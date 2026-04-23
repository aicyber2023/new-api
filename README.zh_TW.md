<div align="center">

![算和雲](/web/public/logo.png)

# 算和雲

🍥 **新一代大模型網關與AI資產管理系統**

<p align="center">
  <strong>繁體中文</strong> |
  <a href="./README.zh_CN.md">简体中文</a> |
  <a href="./README.md">English</a> |
  <a href="./README.fr.md">Français</a> |
  <a href="./README.ja.md">日本語</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/license-AGPL--v3-brightgreen" alt="license">
</p>

<p align="center">
  <a href="#-快速開始">快速開始</a> •
  <a href="#-主要特性">主要特性</a> •
  <a href="#-部署">部署</a> •
  <a href="#-幫助支援">幫助</a>
</p>

</div>

## 📝 專案說明

> [!IMPORTANT]
> - 本專案僅供個人學習使用，不保證穩定性，且不提供任何技術支援
> - 使用者必須在遵循 OpenAI 的 [使用條款](https://openai.com/policies/terms-of-use) 以及**法律法規**的情況下使用，不得用於非法用途
> - 根據 [《生成式人工智慧服務管理暫行辦法》](http://www.cac.gov.cn/2023-07/13/c_1690898327029107.htm) 的要求，請勿對中國地區公眾提供一切未經備案的生成式人工智慧服務

> [!NOTE]
> 本專案基於 [new-api](https://github.com/QuantumNous/new-api) 修改，遵循 AGPL v3 協議。
> 主要新增：**支付寶直連支付**（原生整合，支援沙盒調試）。

---

## 🚀 快速開始

### 環境要求

- Go 1.22+
- Node.js 18+，套件管理器使用 [Bun](https://bun.sh)
- MySQL ≥ 5.7.8 或 PostgreSQL ≥ 9.6（單節點可用 SQLite）

### 編譯與執行

```bash
# 克隆專案
git clone <你的倉庫地址>
cd suanheyun

# 建置前端
cd web && bun install && bun run build && cd ..

# 建置後端
go build -o suanheyun .

# 執行
./suanheyun
```

或使用 Makefile 一鍵建置：

```bash
make all
```

🎉 啟動後訪問 `http://伺服器IP:3000` 即可開始使用！

---

## ✨ 主要特性

### 🎨 核心功能

| 功能 | 說明 |
|------|------|
| 🎨 全新介面 | 現代化使用者介面設計 |
| 🌍 多語言 | 支援簡體中文、繁體中文、英文、法文、日文 |
| 🔄 資料相容 | 完全相容原版 One API 資料庫 |
| 📈 資料看板 | 可視化控制台與統計分析 |
| 🔒 權限管理 | 令牌分組、模型限制、使用者管理 |

### 💰 支付與計費

- ✅ **支付寶直連支付**（原生整合，支援沙盒調試）
- ✅ 線上充值（EPay、Stripe、Waffo、Creem）
- ✅ 按需計費模型定價
- ✅ 快取計費支援（OpenAI、Azure、DeepSeek、Claude、Qwen 等）
- ✅ 靈活的計費策略設定

### 🔐 認證與安全

- 😈 Discord 授權登入
- 🤖 LinuxDO 授權登入
- 📱 Telegram 授權登入
- 🔑 OIDC 統一認證

### 🚀 進階特性

**API 格式支援：**
- ⚡ OpenAI Responses
- ⚡ OpenAI Realtime API（含 Azure）
- ⚡ Claude Messages
- ⚡ Google Gemini
- 🔄 Rerank 模型（Cohere、Jina）

**智慧路由：**
- ⚖️ 渠道加權隨機
- 🔄 失敗自動重試
- 🚦 使用者級模型限速

**格式轉換：**
- 🔄 OpenAI 相容 ⇄ Claude Messages
- 🔄 OpenAI 相容 → Google Gemini
- 🔄 Google Gemini → OpenAI 相容
- 🔄 思維鏈內容轉換

---

## 🚢 部署

### 📋 環境要求

| 元件 | 要求 |
|------|------|
| **資料庫** | SQLite（單節點）/ MySQL ≥ 5.7.8 / PostgreSQL ≥ 9.6 |
| **Go** | 1.22+ |
| **Node.js** | 18+（使用 Bun） |

### ⚙️ 環境變數設定

<details>
<summary>常用環境變數</summary>

| 變數名 | 說明 | 預設值 |
|--------|------|--------|
| `SESSION_SECRET` | 會話金鑰（多節點必填）| - |
| `CRYPTO_SECRET` | 加密金鑰（Redis 必填）| - |
| `SQL_DSN` | 資料庫連接字串 | - |
| `REDIS_CONN_STRING` | Redis 連接字串 | - |
| `STREAMING_TIMEOUT` | 串流逾時（秒）| `300` |
| `AZURE_DEFAULT_API_VERSION` | Azure API 版本 | `2025-04-01-preview` |
| `ERROR_LOG_ENABLED` | 錯誤日誌開關 | `false` |

</details>

### ⚠️ 多節點部署注意事項

> [!WARNING]
> - **必須設定** `SESSION_SECRET`，否則各節點登入狀態不一致
> - **共用 Redis 必須設定** `CRYPTO_SECRET`，否則資料無法解密

---

## 🔗 相關專案

### 上游專案

| 專案 | 說明 |
|------|------|
| [new-api](https://github.com/QuantumNous/new-api) | 直接上游專案（AGPL v3）|
| [One API](https://github.com/songquanpeng/one-api) | 原始專案基礎 |

---

## 💬 幫助支援

歡迎所有形式的貢獻！

- 🐛 回報 Bug
- 💡 提出新功能
- 📝 完善文件
- 🔧 提交程式碼

---

## 📜 開源協議

本專案遵循 [GNU Affero General Public License v3.0 (AGPLv3)](./LICENSE)。

本專案基於 [new-api](https://github.com/QuantumNous/new-api)（AGPL v3）修改，new-api 本身基於 [One API](https://github.com/songquanpeng/one-api)（MIT License）。

---

<div align="center">

### 💖 感謝使用算和雲

</div>
