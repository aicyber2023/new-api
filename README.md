<div align="center">

![算和云](/web/public/logo.png)

# 算和云

🍥 **Next-Generation LLM Gateway and AI Asset Management System**

<p align="center">
  <a href="./README.zh_CN.md">简体中文</a> |
  <a href="./README.zh_TW.md">繁體中文</a> |
  <strong>English</strong> |
  <a href="./README.fr.md">Français</a> |
  <a href="./README.ja.md">日本語</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/license-AGPL--v3-brightgreen" alt="license">
</p>

<p align="center">
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-key-features">Key Features</a> •
  <a href="#-deployment">Deployment</a> •
  <a href="#-help-support">Help</a>
</p>

</div>

## 📝 Project Description

> [!IMPORTANT]
> - This project is for personal learning purposes only, with no guarantee of stability or technical support
> - Users must comply with OpenAI's [Terms of Use](https://openai.com/policies/terms-of-use) and **applicable laws and regulations**, and must not use it for illegal purposes
> - According to the [《Interim Measures for the Management of Generative Artificial Intelligence Services》](http://www.cac.gov.cn/2023-07/13/c_1690898327029107.htm), please do not provide any unregistered generative AI services to the public in China.

> [!NOTE]
> This project is a modified fork of [new-api](https://github.com/QuantumNous/new-api), licensed under AGPL v3.
> Main additions: **Native Alipay direct payment** support.

---

## 🚀 Quick Start

### Requirements

- Go 1.22+
- Node.js 18+ with [Bun](https://bun.sh)
- MySQL ≥ 5.7.8 or PostgreSQL ≥ 9.6 (SQLite supported for single-node)

### Build & Run

```bash
# Clone the project
git clone <your-repo-url>
cd suanheyun

# Build frontend
cd web && bun install && bun run build && cd ..

# Build backend
go build -o suanheyun .

# Run
./suanheyun
```

Or use the all-in-one Makefile:

```bash
make all
```

🎉 After startup, visit `http://your-server-ip:3000` to start using!

---

## ✨ Key Features

### 🎨 Core Functions

| Feature | Description |
|------|------|
| 🎨 New UI | Modern user interface design |
| 🌍 Multi-language | Supports Simplified Chinese, Traditional Chinese, English, French, Japanese |
| 🔄 Data Compatibility | Fully compatible with the original One API database |
| 📈 Data Dashboard | Visual console and statistical analysis |
| 🔒 Permission Management | Token grouping, model restrictions, user management |

### 💰 Payment and Billing

- ✅ **Alipay direct payment** (native integration, sandbox supported)
- ✅ Online recharge (EPay, Stripe, Waffo, Creem)
- ✅ Pay-per-use model pricing
- ✅ Cache billing support (OpenAI, Azure, DeepSeek, Claude, Qwen and all supported models)
- ✅ Flexible billing policy configuration

### 🔐 Authorization and Security

- 😈 Discord authorization login
- 🤖 LinuxDO authorization login
- 📱 Telegram authorization login
- 🔑 OIDC unified authentication

### 🚀 Advanced Features

**API Format Support:**
- ⚡ OpenAI Responses
- ⚡ OpenAI Realtime API (including Azure)
- ⚡ Claude Messages
- ⚡ Google Gemini
- 🔄 Rerank Models (Cohere, Jina)

**Intelligent Routing:**
- ⚖️ Channel weighted random
- 🔄 Automatic retry on failure
- 🚦 User-level model rate limiting

**Format Conversion:**
- 🔄 OpenAI Compatible ⇄ Claude Messages
- 🔄 OpenAI Compatible → Google Gemini
- 🔄 Google Gemini → OpenAI Compatible
- 🔄 Thinking-to-content functionality

---

## 🤖 Model Support

| Model Type | Description |
|---------|------|
| 🤖 OpenAI-Compatible | OpenAI compatible models |
| 🤖 OpenAI Responses | OpenAI Responses format |
| 🎨 Midjourney-Proxy | Midjourney-Proxy(Plus) |
| 🎵 Suno-API | Suno API |
| 🔄 Rerank | Cohere, Jina |
| 💬 Claude | Messages format |
| 🌐 Gemini | Google Gemini format |
| 🔧 Dify | ChatFlow mode |
| 🎯 Custom | Supports complete call address |

---

## 🚢 Deployment

### 📋 Requirements

| Component | Requirement |
|------|------|
| **Database** | SQLite (single-node) / MySQL ≥ 5.7.8 / PostgreSQL ≥ 9.6 |
| **Go** | 1.22+ |
| **Node.js** | 18+ (with Bun) |

### ⚙️ Environment Variables

<details>
<summary>Common environment variables</summary>

| Variable | Description | Default |
|--------|------|--------|
| `SESSION_SECRET` | Session secret (required for multi-node) | - |
| `CRYPTO_SECRET` | Encryption secret (required with Redis) | - |
| `SQL_DSN` | Database connection string | - |
| `REDIS_CONN_STRING` | Redis connection string | - |
| `STREAMING_TIMEOUT` | Streaming timeout (seconds) | `300` |
| `AZURE_DEFAULT_API_VERSION` | Azure API version | `2025-04-01-preview` |
| `ERROR_LOG_ENABLED` | Error log switch | `false` |

</details>

### ⚠️ Multi-node Deployment

> [!WARNING]
> - **Must set** `SESSION_SECRET` — otherwise login state is inconsistent across nodes
> - **Shared Redis must set** `CRYPTO_SECRET` — otherwise data cannot be decrypted

---

## 🔗 Related Projects

### Upstream Projects

| Project | Description |
|------|------|
| [new-api](https://github.com/QuantumNous/new-api) | Direct upstream project (AGPL v3) |
| [One API](https://github.com/songquanpeng/one-api) | Original project base |

---

## 💬 Help & Support

Welcome all forms of contribution!

- 🐛 Report Bugs
- 💡 Propose New Features
- 📝 Improve Documentation
- 🔧 Submit Code

---

## 📜 License

This project is licensed under the [GNU Affero General Public License v3.0 (AGPLv3)](./LICENSE).

This is a modified fork of [new-api](https://github.com/QuantumNous/new-api) (AGPL v3), which is itself based on [One API](https://github.com/songquanpeng/one-api) (MIT License).

---

<div align="center">

### 💖 Thank you for using 算和云

</div>
