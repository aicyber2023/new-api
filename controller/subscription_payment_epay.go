package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/QuantumNous/new-api/setting/system_setting"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
)

type SubscriptionEpayPayRequest struct {
	PlanId        int    `json:"plan_id"`
	PaymentMethod string `json:"payment_method"`
}

func SubscriptionRequestEpay(c *gin.Context) {
	var req SubscriptionEpayPayRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.PlanId <= 0 {
		common.ApiErrorMsg(c, "参数错误")
		return
	}

	plan, err := model.GetSubscriptionPlanById(req.PlanId)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	if !plan.Enabled {
		common.ApiErrorMsg(c, "套餐未启用")
		return
	}
	if plan.PriceAmount < 0.01 {
		common.ApiErrorMsg(c, "套餐金额过低")
		return
	}
	if !operation_setting.ContainsPayMethod(req.PaymentMethod) {
		common.ApiErrorMsg(c, "支付方式不存在")
		return
	}

	userId := c.GetInt("id")
	if plan.MaxPurchasePerUser > 0 {
		count, err := model.CountUserSubscriptionsByPlan(userId, plan.Id)
		if err != nil {
			common.ApiError(c, err)
			return
		}
		if count >= int64(plan.MaxPurchasePerUser) {
			common.ApiErrorMsg(c, "已达到该套餐购买上限")
			return
		}
	}

	client, err := GetAlipayClient()
	if err != nil {
		common.ApiErrorMsg(c, "当前管理员未配置支付信息")
		return
	}

	callBackAddress := service.GetCallbackAddress()
	notifyUrl := callBackAddress + "/api/subscription/epay/notify"
	returnUrl := system_setting.ServerAddress + "/console/log"

	tradeNo := fmt.Sprintf("%s%d", common.GetRandomString(6), time.Now().Unix())
	tradeNo = fmt.Sprintf("SUBUSR%dNO%s", userId, tradeNo)

	order := &model.SubscriptionOrder{
		UserId:        userId,
		PlanId:        plan.Id,
		Money:         plan.PriceAmount,
		TradeNo:       tradeNo,
		PaymentMethod: req.PaymentMethod,
		CreateTime:    time.Now().Unix(),
		Status:        common.TopUpStatusPending,
	}
	if err := order.Insert(); err != nil {
		common.ApiErrorMsg(c, "创建订单失败")
		return
	}

	bm := make(gopay.BodyMap)
	bm.Set("subject", fmt.Sprintf("订阅套餐：%s", plan.Title))
	bm.Set("out_trade_no", tradeNo)
	bm.Set("total_amount", strconv.FormatFloat(plan.PriceAmount, 'f', 2, 64))
	bm.Set("return_url", returnUrl)
	bm.Set("notify_url", notifyUrl)

	payUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		_ = model.ExpireSubscriptionOrder(tradeNo)
		log.Printf("订阅套餐拉起支付宝支付失败: %v", err)
		common.ApiErrorMsg(c, "拉起支付失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "url": payUrl})
}

func SubscriptionEpayNotify(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		log.Println("订阅回调解析失败:", err)
		_, _ = c.Writer.Write([]byte("fail"))
		return
	}

	params := make(gopay.BodyMap)
	for k, vs := range c.Request.Form {
		if len(vs) > 0 {
			params.Set(k, vs[0])
		}
	}

	if len(params) == 0 {
		_, _ = c.Writer.Write([]byte("fail"))
		return
	}

	ok, err := alipay.VerifySign(operation_setting.AlipayPublicKey, params)
	if err != nil || !ok {
		log.Println("订阅回调验签失败:", err)
		_, _ = c.Writer.Write([]byte("fail"))
		return
	}

	tradeStatus := params.GetString("trade_status")
	if tradeStatus != "TRADE_SUCCESS" && tradeStatus != "TRADE_FINISHED" {
		_, _ = c.Writer.Write([]byte("success"))
		return
	}

	outTradeNo := params.GetString("out_trade_no")
	if outTradeNo == "" {
		_, _ = c.Writer.Write([]byte("fail"))
		return
	}

	LockOrder(outTradeNo)
	defer UnlockOrder(outTradeNo)

	if err := model.CompleteSubscriptionOrder(outTradeNo, params.GetString("trade_no")); err != nil {
		log.Printf("订阅回调完成订单失败: %v", err)
		_, _ = c.Writer.Write([]byte("fail"))
		return
	}

	_, _ = c.Writer.Write([]byte("success"))
}

func SubscriptionEpayReturn(c *gin.Context) {
	c.Redirect(http.StatusFound, system_setting.ServerAddress+"/console/topup?pay=success")
}
