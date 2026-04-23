<div align="center">

![算和云](/web/public/logo.png)

# 算和云 (SuanHeYun)

🍥 **次世代大規模言語モデルゲートウェイとAI資産管理システム**

<p align="center">
  <a href="./README.zh_CN.md">简体中文</a> |
  <a href="./README.zh_TW.md">繁體中文</a> |
  <a href="./README.md">English</a> |
  <a href="./README.fr.md">Français</a> |
  <strong>日本語</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/license-AGPL--v3-brightgreen" alt="license">
</p>

<p align="center">
  <a href="#-クイックスタート">クイックスタート</a> •
  <a href="#-主な機能">主な機能</a> •
  <a href="#-デプロイ">デプロイ</a> •
  <a href="#-ヘルプサポート">ヘルプ</a>
</p>

</div>

## 📝 プロジェクト説明

> [!IMPORTANT]
> - このプロジェクトは個人学習目的のみで、安定性の保証や技術サポートは提供されません
> - ユーザーはOpenAIの[利用規約](https://openai.com/policies/terms-of-use)および**適用法令**を遵守し、違法な目的に使用しないでください
> - 中国の生成AI規制に基づき、未登録の生成AIサービスを中国の一般公衆に提供しないでください

> [!NOTE]
> このプロジェクトは [new-api](https://github.com/QuantumNous/new-api) をベースに改変されており、AGPL v3 ライセンスに従います。
> 主な追加機能：**Alipay ネイティブ直接決済**（サンドボックス対応）。

---

## 🚀 クイックスタート

### 動作要件

- Go 1.22+
- Node.js 18+、パッケージマネージャーは [Bun](https://bun.sh) を使用
- MySQL ≥ 5.7.8 または PostgreSQL ≥ 9.6（単一ノードは SQLite 可）

### ビルドと起動

```bash
# プロジェクトをクローン
git clone <あなたのリポジトリURL>
cd suanheyun

# フロントエンドをビルド
cd web && bun install && bun run build && cd ..

# バックエンドをビルド
go build -o suanheyun .

# 起動
./suanheyun
```

または Makefile で一括ビルド：

```bash
make all
```

🎉 起動後、`http://サーバーIP:3000` にアクセスして使用開始！

---

## ✨ 主な機能

### 🎨 コア機能

| 機能 | 説明 |
|------|------|
| 🎨 新しいUI | モダンなユーザーインターフェースデザイン |
| 🌍 多言語 | 簡体字中国語、繁体字中国語、英語、フランス語、日本語をサポート |
| 🔄 データ互換性 | 元の One API データベースと完全互換 |
| 📈 データダッシュボード | ビジュアルコンソールと統計分析 |
| 🔒 権限管理 | トークングループ化、モデル制限、ユーザー管理 |

### 💰 支払いと請求

- ✅ **Alipay ネイティブ直接決済**（サンドボックス対応）
- ✅ オンラインチャージ（EPay、Stripe、Waffo、Creem）
- ✅ 従量課金モデル価格設定
- ✅ キャッシュ請求サポート（OpenAI、Azure、DeepSeek、Claude、Qwen など）
- ✅ 柔軟な請求ポリシー設定

### 🔐 認証とセキュリティ

- 😈 Discord 認証ログイン
- 🤖 LinuxDO 認証ログイン
- 📱 Telegram 認証ログイン
- 🔑 OIDC 統合認証

### 🚀 高度な機能

**API フォーマットサポート：**
- ⚡ OpenAI Responses
- ⚡ OpenAI Realtime API（Azure 含む）
- ⚡ Claude Messages
- ⚡ Google Gemini
- 🔄 Rerank モデル（Cohere、Jina）

**インテリジェントルーティング：**
- ⚖️ チャンネル重み付きランダム
- 🔄 失敗時の自動リトライ
- 🚦 ユーザーレベルモデルレート制限

**フォーマット変換：**
- 🔄 OpenAI 互換 ⇄ Claude Messages
- 🔄 OpenAI 互換 → Google Gemini
- 🔄 Google Gemini → OpenAI 互換
- 🔄 思考連鎖コンテンツ変換

---

## 🚢 デプロイ

### 📋 動作要件

| コンポーネント | 要件 |
|------|------|
| **データベース** | SQLite（単一ノード）/ MySQL ≥ 5.7.8 / PostgreSQL ≥ 9.6 |
| **Go** | 1.22+ |
| **Node.js** | 18+（Bun 使用） |

### ⚙️ 環境変数設定

<details>
<summary>主要な環境変数</summary>

| 変数名 | 説明 | デフォルト値 |
|--------|------|--------|
| `SESSION_SECRET` | セッションシークレット（マルチノード必須）| - |
| `CRYPTO_SECRET` | 暗号化シークレット（Redis 必須）| - |
| `SQL_DSN` | データベース接続文字列 | - |
| `REDIS_CONN_STRING` | Redis 接続文字列 | - |
| `STREAMING_TIMEOUT` | ストリーミングタイムアウト（秒）| `300` |
| `AZURE_DEFAULT_API_VERSION` | Azure API バージョン | `2025-04-01-preview` |
| `ERROR_LOG_ENABLED` | エラーログスイッチ | `false` |

</details>

### ⚠️ マルチノードデプロイ

> [!WARNING]
> - **必ず** `SESSION_SECRET` を設定してください — 設定しないとノード間でログイン状態が不整合になります
> - **共有 Redis には必ず** `CRYPTO_SECRET` を設定してください — 設定しないとデータを復号できません

---

## 🔗 関連プロジェクト

### アップストリームプロジェクト

| プロジェクト | 説明 |
|------|------|
| [new-api](https://github.com/QuantumNous/new-api) | 直接アップストリーム（AGPL v3）|
| [One API](https://github.com/songquanpeng/one-api) | 元のプロジェクトベース |

---

## 💬 ヘルプサポート

あらゆる形の貢献を歓迎します！

- 🐛 バグ報告
- 💡 新機能の提案
- 📝 ドキュメントの改善
- 🔧 コードの提出

---

## 📜 ライセンス

このプロジェクトは [GNU Affero General Public License v3.0 (AGPLv3)](./LICENSE) の下でライセンスされています。

このプロジェクトは [new-api](https://github.com/QuantumNous/new-api)（AGPL v3）のフォークであり、new-api 自体は [One API](https://github.com/songquanpeng/one-api)（MIT License）をベースにしています。

---

<div align="center">

### 💖 算和云 (SuanHeYun) をご利用いただきありがとうございます

</div>
