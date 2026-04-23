<div align="center">

![算和云](/web/public/logo.png)

# 算和云

🍥 **新一代大模型网关与AI资产管理系统**

<p align="center">
  <strong>简体中文</strong> |
  <a href="./README.zh_TW.md">繁體中文</a> |
  <a href="./README.md">English</a> |
  <a href="./README.fr.md">Français</a> |
  <a href="./README.ja.md">日本語</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/license-AGPL--v3-brightgreen" alt="license">
</p>

<p align="center">
  <a href="#-快速开始">快速开始</a> •
  <a href="#-主要特性">主要特性</a> •
  <a href="#-部署">部署</a> •
  <a href="#-帮助支持">帮助</a>
</p>

</div>

## 📝 项目说明

> [!IMPORTANT]
> - 本项目仅供个人学习使用，不保证稳定性，且不提供任何技术支持
> - 使用者必须在遵循 OpenAI 的 [使用条款](https://openai.com/policies/terms-of-use) 以及**法律法规**的情况下使用，不得用于非法用途
> - 根据 [《生成式人工智能服务管理暂行办法》](http://www.cac.gov.cn/2023-07/13/c_1690898327029107.htm) 的要求，请勿对中国地区公众提供一切未经备案的生成式人工智能服务

> [!NOTE]
> 本项目基于 [new-api](https://github.com/QuantumNous/new-api) 修改，遵循 AGPL v3 协议。
> 主要新增：**支付宝直连支付**（原生集成，支持沙箱调试）。

---

## 🚀 快速开始

### 环境要求

- Go 1.22+
- Node.js 18+，包管理器使用 [Bun](https://bun.sh)
- MySQL ≥ 5.7.8 或 PostgreSQL ≥ 9.6（单节点可用 SQLite）

### 编译与运行

```bash
# 克隆项目
git clone <你的仓库地址>
cd suanheyun

# 构建前端
cd web && bun install && bun run build && cd ..

# 构建后端
go build -o suanheyun .

# 运行
./suanheyun
```

或使用 Makefile 一键构建：

```bash
make all
```

🎉 启动后访问 `http://服务器IP:3000` 即可开始使用！

---

## ✨ 主要特性

### 🎨 核心功能

| 功能 | 说明 |
|------|------|
| 🎨 全新界面 | 现代化用户界面设计 |
| 🌍 多语言 | 支持简体中文、繁体中文、英文、法文、日文 |
| 🔄 数据兼容 | 完全兼容原版 One API 数据库 |
| 📈 数据看板 | 可视化控制台与统计分析 |
| 🔒 权限管理 | 令牌分组、模型限制、用户管理 |

### 💰 支付与计费

- ✅ **支付宝直连支付**（原生集成，支持沙箱调试）
- ✅ 在线充值（EPay、Stripe、Waffo、Creem）
- ✅ 按需计费模型定价
- ✅ 缓存计费支持（OpenAI、Azure、DeepSeek、Claude、Qwen 等全部支持模型）
- ✅ 灵活的计费策略配置

### 🔐 认证与安全

- 😈 Discord 授权登录
- 🤖 LinuxDO 授权登录
- 📱 Telegram 授权登录
- 🔑 OIDC 统一认证

### 🚀 进阶特性

**API 格式支持：**
- ⚡ OpenAI Responses
- ⚡ OpenAI Realtime API（含 Azure）
- ⚡ Claude Messages
- ⚡ Google Gemini
- 🔄 Rerank 模型（Cohere、Jina）

**智能路由：**
- ⚖️ 渠道加权随机
- 🔄 失败自动重试
- 🚦 用户级模型限速

**格式转换：**
- 🔄 OpenAI 兼容 ⇄ Claude Messages
- 🔄 OpenAI 兼容 → Google Gemini
- 🔄 Google Gemini → OpenAI 兼容
- 🔄 思维链内容转换

---

## 🤖 模型支持

| 模型类型 | 说明 |
|---------|------|
| 🤖 OpenAI 兼容 | OpenAI 兼容模型 |
| 🤖 OpenAI Responses | OpenAI Responses 格式 |
| 🎨 Midjourney-Proxy | Midjourney-Proxy(Plus) |
| 🎵 Suno-API | Suno API |
| 🔄 Rerank | Cohere、Jina |
| 💬 Claude | Messages 格式 |
| 🌐 Gemini | Google Gemini 格式 |
| 🔧 Dify | ChatFlow 模式 |
| 🎯 自定义 | 支持完整调用地址 |

---

## 🚢 部署

### 📋 环境要求

| 组件 | 要求 |
|------|------|
| **数据库** | SQLite（单节点）/ MySQL ≥ 5.7.8 / PostgreSQL ≥ 9.6 |
| **Go** | 1.22+ |
| **Node.js** | 18+（使用 Bun） |

### ⚙️ 环境变量配置

<details>
<summary>常用环境变量</summary>

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| `SESSION_SECRET` | 会话密钥（多节点必填）| - |
| `CRYPTO_SECRET` | 加密密钥（Redis 必填）| - |
| `SQL_DSN` | 数据库连接字符串 | - |
| `REDIS_CONN_STRING` | Redis 连接字符串 | - |
| `STREAMING_TIMEOUT` | 流式超时（秒）| `300` |
| `AZURE_DEFAULT_API_VERSION` | Azure API 版本 | `2025-04-01-preview` |
| `ERROR_LOG_ENABLED` | 错误日志开关 | `false` |

</details>

### ⚠️ 多节点部署注意事项

> [!WARNING]
> - **必须设置** `SESSION_SECRET`，否则各节点登录状态不一致
> - **共享 Redis 必须设置** `CRYPTO_SECRET`，否则数据无法解密

---

## 🔗 相关项目

### 上游项目

| 项目 | 说明 |
|------|------|
| [new-api](https://github.com/QuantumNous/new-api) | 直接上游项目（AGPL v3）|
| [One API](https://github.com/songquanpeng/one-api) | 原始项目基础 |

---

## 💬 帮助支持

欢迎所有形式的贡献！

- 🐛 反馈 Bug
- 💡 提出新功能
- 📝 完善文档
- 🔧 提交代码

---

## 📜 开源协议

本项目遵循 [GNU Affero General Public License v3.0 (AGPLv3)](./LICENSE)。

本项目基于 [new-api](https://github.com/QuantumNous/new-api)（AGPL v3）修改，new-api 本身基于 [One API](https://github.com/songquanpeng/one-api)（MIT License）。

---

<div align="center">

### 💖 感谢使用算和云

</div>
