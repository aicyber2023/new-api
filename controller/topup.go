package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/setting"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/QuantumNous/new-api/setting/system_setting"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/shopspring/decimal"
)

func GetTopUpInfo(c *gin.Context) {
	payMethods := operation_setting.PayMethods

	if setting.StripeApiSecret != "" && setting.StripeWebhookSecret != "" && setting.StripePriceId != "" {
		hasStripe := false
		for _, method := range payMethods {
			if method["type"] == "stripe" {
				hasStripe = true
				break
			}
		}
		if !hasStripe {
			stripeMethod := map[string]string{
				"name":      "Stripe",
				"type":      "stripe",
				"color":     "rgba(var(--semi-purple-5), 1)",
				"min_topup": strconv.Itoa(setting.StripeMinTopUp),
			}
			payMethods = append(payMethods, stripeMethod)
		}
	}

	enableWaffo := setting.WaffoEnabled &&
		((!setting.WaffoSandbox &&
			setting.WaffoApiKey != "" &&
			setting.WaffoPrivateKey != "" &&
			setting.WaffoPublicCert != "") ||
			(setting.WaffoSandbox &&
				setting.WaffoSandboxApiKey != "" &&
				setting.WaffoSandboxPrivateKey != "" &&
				setting.WaffoSandboxPublicCert != ""))
	if enableWaffo {
		hasWaffo := false
		for _, method := range payMethods {
			if method["type"] == "waffo" {
				hasWaffo = true
				break
			}
		}
		if !hasWaffo {
			waffoMethod := map[string]string{
				"name":      "Waffo (Global Payment)",
				"type":      "waffo",
				"color":     "rgba(var(--semi-blue-5), 1)",
				"min_topup": strconv.Itoa(setting.WaffoMinTopUp),
			}
			payMethods = append(payMethods, waffoMethod)
		}
	}

	data := gin.H{
		"enable_online_topup": operation_setting.AlipayAppId != "" && operation_setting.AlipayPrivateKey != "",
		"enable_stripe_topup": setting.StripeApiSecret != "" && setting.StripeWebhookSecret != "" && setting.StripePriceId != "",
		"enable_creem_topup":  setting.CreemApiKey != "" && setting.CreemProducts != "[]",
		"enable_waffo_topup":  enableWaffo,
		"waffo_pay_methods": func() interface{} {
			if enableWaffo {
				return setting.GetWaffoPayMethods()
			}
			return nil
		}(),
		"creem_products":   setting.CreemProducts,
		"pay_methods":      payMethods,
		"min_topup":        operation_setting.MinTopUp,
		"stripe_min_topup": setting.StripeMinTopUp,
		"waffo_min_topup":  setting.WaffoMinTopUp,
		"amount_options":   operation_setting.GetPaymentSetting().AmountOptions,
		"discount":         operation_setting.GetPaymentSetting().AmountDiscount,
	}
	common.ApiSuccess(c, data)
}

type EpayRequest struct {
	Amount        int64  `json:"amount"`
	PaymentMethod string `json:"payment_method"`
}

type AmountRequest struct {
	Amount int64 `json:"amount"`
}

func GetAlipayClient() (*alipay.Client, error) {
	if operation_setting.AlipayAppId == "" || operation_setting.AlipayPrivateKey == "" {
		return nil, errors.New("支付宝未配置")
	}
	isProd := !operation_setting.AlipaySandbox
	client, err := alipay.NewClient(
		operation_setting.AlipayAppId,
		operation_setting.AlipayPrivateKey,
		isProd,
	)
	if err != nil {
		return nil, err
	}
	if operation_setting.AlipayPublicKey != "" {
		pubKey := operation_setting.AlipayPublicKey
		if !strings.Contains(pubKey, "-----BEGIN") {
			pubKey = "-----BEGIN PUBLIC KEY-----\n" + pubKey + "\n-----END PUBLIC KEY-----"
		}
		client.AutoVerifySign([]byte(pubKey))
	}
	return client, nil
}

func getPayMoney(amount int64, group string) float64 {
	dAmount := decimal.NewFromInt(amount)
	if operation_setting.GetQuotaDisplayType() == operation_setting.QuotaDisplayTypeTokens {
		dQuotaPerUnit := decimal.NewFromFloat(common.QuotaPerUnit)
		dAmount = dAmount.Div(dQuotaPerUnit)
	}

	topupGroupRatio := common.GetTopupGroupRatio(group)
	if topupGroupRatio == 0 {
		topupGroupRatio = 1
	}

	dTopupGroupRatio := decimal.NewFromFloat(topupGroupRatio)
	dPrice := decimal.NewFromFloat(operation_setting.Price)
	discount := 1.0
	if ds, ok := operation_setting.GetPaymentSetting().AmountDiscount[int(amount)]; ok {
		if ds > 0 {
			discount = ds
		}
	}
	dDiscount := decimal.NewFromFloat(discount)

	payMoney := dAmount.Mul(dPrice).Mul(dTopupGroupRatio).Mul(dDiscount)
	return payMoney.InexactFloat64()
}

func getMinTopup() int64 {
	minTopup := operation_setting.MinTopUp
	if operation_setting.GetQuotaDisplayType() == operation_setting.QuotaDisplayTypeTokens {
		dMinTopup := decimal.NewFromInt(int64(minTopup))
		dQuotaPerUnit := decimal.NewFromFloat(common.QuotaPerUnit)
		minTopup = int(dMinTopup.Mul(dQuotaPerUnit).IntPart())
	}
	return int64(minTopup)
}

func RequestEpay(c *gin.Context) {
	var req EpayRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"message": "error", "data": "参数错误"})
		return
	}
	if req.Amount < getMinTopup() {
		c.JSON(200, gin.H{"message": "error", "data": fmt.Sprintf("充值数量不能小于 %d", getMinTopup())})
		return
	}

	id := c.GetInt("id")
	group, err := model.GetUserGroup(id, true)
	if err != nil {
		c.JSON(200, gin.H{"message": "error", "data": "获取用户分组失败"})
		return
	}
	payMoney := getPayMoney(req.Amount, group)
	if payMoney < 0.01 {
		c.JSON(200, gin.H{"message": "error", "data": "充值金额过低"})
		return
	}

	if !operation_setting.ContainsPayMethod(req.PaymentMethod) {
		c.JSON(200, gin.H{"message": "error", "data": "支付方式不存在"})
		return
	}

	client, err := GetAlipayClient()
	if err != nil {
		c.JSON(200, gin.H{"message": "error", "data": "当前管理员未配置支付信息"})
		return
	}

	callBackAddress := service.GetCallbackAddress()
	notifyUrl := callBackAddress + "/api/user/alipay/notify"
	returnUrl := system_setting.ServerAddress + "/console/log"

	tradeNo := fmt.Sprintf("%s%d", common.GetRandomString(6), time.Now().Unix())
	tradeNo = fmt.Sprintf("USR%dNO%s", id, tradeNo)

	amount := req.Amount
	if operation_setting.GetQuotaDisplayType() == operation_setting.QuotaDisplayTypeTokens {
		dAmount := decimal.NewFromInt(int64(amount))
		dQuotaPerUnit := decimal.NewFromFloat(common.QuotaPerUnit)
		amount = dAmount.Div(dQuotaPerUnit).IntPart()
	}

	bm := make(gopay.BodyMap)
	bm.Set("subject", fmt.Sprintf("充值 %d 额度", amount))
	bm.Set("out_trade_no", tradeNo)
	bm.Set("total_amount", strconv.FormatFloat(payMoney, 'f', 2, 64))
	bm.Set("return_url", returnUrl)
	bm.Set("notify_url", notifyUrl)

	payUrl, err := client.TradePagePay(c.Request.Context(), bm)
	if err != nil {
		log.Printf("支付宝拉起支付失败: %v", err)
		c.JSON(200, gin.H{"message": "error", "data": "拉起支付失败"})
		return
	}

	topUp := &model.TopUp{
		UserId:        id,
		Amount:        amount,
		Money:         payMoney,
		TradeNo:       tradeNo,
		PaymentMethod: req.PaymentMethod,
		CreateTime:    time.Now().Unix(),
		Status:        "pending",
	}
	if err = topUp.Insert(); err != nil {
		c.JSON(200, gin.H{"message": "error", "data": "创建订单失败"})
		return
	}

	c.JSON(200, gin.H{"message": "success", "url": payUrl, "data": gin.H{}})
}

// tradeNo lock
var orderLocks sync.Map
var createLock sync.Mutex

type refCountedMutex struct {
	mu       sync.Mutex
	refCount int
}

func LockOrder(tradeNo string) {
	createLock.Lock()
	var rcm *refCountedMutex
	if v, ok := orderLocks.Load(tradeNo); ok {
		rcm = v.(*refCountedMutex)
	} else {
		rcm = &refCountedMutex{}
		orderLocks.Store(tradeNo, rcm)
	}
	rcm.refCount++
	createLock.Unlock()
	rcm.mu.Lock()
}

func UnlockOrder(tradeNo string) {
	v, ok := orderLocks.Load(tradeNo)
	if !ok {
		return
	}
	rcm := v.(*refCountedMutex)
	rcm.mu.Unlock()

	createLock.Lock()
	rcm.refCount--
	if rcm.refCount == 0 {
		orderLocks.Delete(tradeNo)
	}
	createLock.Unlock()
}

func EpayNotify(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		log.Println("支付宝回调解析失败:", err)
		c.String(http.StatusOK, "fail")
		return
	}

	// 收集所有表单参数
	params := make(gopay.BodyMap)
	for k, vs := range c.Request.Form {
		if len(vs) > 0 {
			params.Set(k, vs[0])
		}
	}

	// 验签
	pubKey := operation_setting.AlipayPublicKey
	if !strings.Contains(pubKey, "-----BEGIN") {
		pubKey = "-----BEGIN PUBLIC KEY-----\n" + pubKey + "\n-----END PUBLIC KEY-----"
	}
	ok, err := alipay.VerifySign(pubKey, params)
	if err != nil || !ok {
		log.Println("支付宝回调验签失败:", err)
		c.String(http.StatusOK, "fail")
		return
	}

	tradeStatus := params.GetString("trade_status")
	if tradeStatus != "TRADE_SUCCESS" && tradeStatus != "TRADE_FINISHED" {
		log.Printf("支付宝回调非成功状态: %s", tradeStatus)
		c.String(http.StatusOK, "success") // 告知支付宝已收到，不重试
		return
	}

	outTradeNo := params.GetString("out_trade_no")
	if outTradeNo == "" {
		log.Println("支付宝回调缺少 out_trade_no")
		c.String(http.StatusOK, "fail")
		return
	}

	LockOrder(outTradeNo)
	defer UnlockOrder(outTradeNo)

	topUp := model.GetTopUpByTradeNo(outTradeNo)
	if topUp == nil {
		log.Printf("支付宝回调未找到订单: %s", outTradeNo)
		c.String(http.StatusOK, "fail")
		return
	}

	if topUp.Status == "pending" {
		topUp.Status = "success"
		if err := topUp.Update(); err != nil {
			log.Printf("支付宝回调更新订单失败: %v", topUp)
			c.String(http.StatusOK, "fail")
			return
		}
		dAmount := decimal.NewFromInt(int64(topUp.Amount))
		dQuotaPerUnit := decimal.NewFromFloat(common.QuotaPerUnit)
		quotaToAdd := int(dAmount.Mul(dQuotaPerUnit).IntPart())
		if err := model.IncreaseUserQuota(topUp.UserId, quotaToAdd, true); err != nil {
			log.Printf("支付宝回调更新用户额度失败: %v", topUp)
			c.String(http.StatusOK, "fail")
			return
		}
		log.Printf("支付宝充值成功 %v", topUp)
		model.RecordLog(topUp.UserId, model.LogTypeTopup, fmt.Sprintf("使用支付宝充值成功，充值金额: %v，支付金额：%.2f", logger.LogQuota(quotaToAdd), topUp.Money))
	}

	c.String(http.StatusOK, "success")
}

func RequestAmount(c *gin.Context) {
	var req AmountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"message": "error", "data": "参数错误"})
		return
	}
	if req.Amount < getMinTopup() {
		c.JSON(200, gin.H{"message": "error", "data": fmt.Sprintf("充值数量不能小于 %d", getMinTopup())})
		return
	}
	id := c.GetInt("id")
	group, err := model.GetUserGroup(id, true)
	if err != nil {
		c.JSON(200, gin.H{"message": "error", "data": "获取用户分组失败"})
		return
	}
	payMoney := getPayMoney(req.Amount, group)
	if payMoney <= 0.01 {
		c.JSON(200, gin.H{"message": "error", "data": "充值金额过低"})
		return
	}
	c.JSON(200, gin.H{"message": "success", "data": strconv.FormatFloat(payMoney, 'f', 2, 64)})
}

func GetUserTopUps(c *gin.Context) {
	userId := c.GetInt("id")
	pageInfo := common.GetPageQuery(c)
	keyword := c.Query("keyword")

	var (
		topups []*model.TopUp
		total  int64
		err    error
	)
	if keyword != "" {
		topups, total, err = model.SearchUserTopUps(userId, keyword, pageInfo)
	} else {
		topups, total, err = model.GetUserTopUps(userId, pageInfo)
	}
	if err != nil {
		common.ApiError(c, err)
		return
	}

	pageInfo.SetTotal(int(total))
	pageInfo.SetItems(topups)
	common.ApiSuccess(c, pageInfo)
}

func GetAllTopUps(c *gin.Context) {
	pageInfo := common.GetPageQuery(c)
	keyword := c.Query("keyword")

	var (
		topups []*model.TopUp
		total  int64
		err    error
	)
	if keyword != "" {
		topups, total, err = model.SearchAllTopUps(keyword, pageInfo)
	} else {
		topups, total, err = model.GetAllTopUps(pageInfo)
	}
	if err != nil {
		common.ApiError(c, err)
		return
	}

	pageInfo.SetTotal(int(total))
	pageInfo.SetItems(topups)
	common.ApiSuccess(c, pageInfo)
}

type AdminCompleteTopupRequest struct {
	TradeNo string `json:"trade_no"`
}

func AdminCompleteTopUp(c *gin.Context) {
	var req AdminCompleteTopupRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.TradeNo == "" {
		common.ApiErrorMsg(c, "参数错误")
		return
	}

	LockOrder(req.TradeNo)
	defer UnlockOrder(req.TradeNo)

	if err := model.ManualCompleteTopUp(req.TradeNo); err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, nil)
}
