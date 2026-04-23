<div align="center">

![算和云](/web/public/logo.png)

# 算和云 (SuanHeYun)

🍥 **Passerelle de modèles LLM de nouvelle génération et système de gestion d'actifs IA**

<p align="center">
  <a href="./README.zh_CN.md">简体中文</a> |
  <a href="./README.zh_TW.md">繁體中文</a> |
  <a href="./README.md">English</a> |
  <strong>Français</strong> |
  <a href="./README.ja.md">日本語</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/license-AGPL--v3-brightgreen" alt="licence">
</p>

<p align="center">
  <a href="#-démarrage-rapide">Démarrage rapide</a> •
  <a href="#-fonctionnalités-clés">Fonctionnalités</a> •
  <a href="#-déploiement">Déploiement</a> •
  <a href="#-aide-support">Aide</a>
</p>

</div>

## 📝 Description du projet

> [!IMPORTANT]
> - Ce projet est destiné à un usage personnel uniquement, sans garantie de stabilité ni support technique
> - Les utilisateurs doivent respecter les [Conditions d'utilisation](https://openai.com/policies/terms-of-use) d'OpenAI ainsi que les **lois et réglementations applicables**
> - Conformément aux réglementations chinoises sur l'IA générative, ne pas fournir de services d'IA non enregistrés au public en Chine

> [!NOTE]
> Ce projet est un fork modifié de [new-api](https://github.com/QuantumNous/new-api), sous licence AGPL v3.
> Principal ajout : **Paiement Alipay natif** (intégration directe, bac à sable supporté).

---

## 🚀 Démarrage rapide

### Prérequis

- Go 1.22+
- Node.js 18+ avec [Bun](https://bun.sh)
- MySQL ≥ 5.7.8 ou PostgreSQL ≥ 9.6 (SQLite supporté pour nœud unique)

### Compilation et démarrage

```bash
# Cloner le projet
git clone <url-de-votre-dépôt>
cd suanheyun

# Construire le frontend
cd web && bun install && bun run build && cd ..

# Construire le backend
go build -o suanheyun .

# Démarrer
./suanheyun
```

Ou avec le Makefile tout-en-un :

```bash
make all
```

🎉 Après le démarrage, visitez `http://ip-serveur:3000` pour commencer !

---

## ✨ Fonctionnalités clés

### 🎨 Fonctions principales

| Fonctionnalité | Description |
|------|------|
| 🎨 Nouvelle interface | Design d'interface utilisateur moderne |
| 🌍 Multi-langue | Chinois simplifié, traditionnel, anglais, français, japonais |
| 🔄 Compatibilité des données | Entièrement compatible avec la base de données One API originale |
| 📈 Tableau de bord | Console visuelle et analyse statistique |
| 🔒 Gestion des permissions | Groupes de tokens, restrictions de modèles, gestion des utilisateurs |

### 💰 Paiement et facturation

- ✅ **Paiement Alipay natif** (intégration directe, bac à sable supporté)
- ✅ Recharge en ligne (EPay, Stripe, Waffo, Creem)
- ✅ Tarification à l'usage par modèle
- ✅ Facturation avec cache (OpenAI, Azure, DeepSeek, Claude, Qwen et tous les modèles supportés)
- ✅ Configuration flexible des politiques de facturation

### 🔐 Authentification et sécurité

- 😈 Connexion via Discord
- 🤖 Connexion via LinuxDO
- 📱 Connexion via Telegram
- 🔑 Authentification unifiée OIDC

### 🚀 Fonctionnalités avancées

**Formats API supportés :**
- ⚡ OpenAI Responses
- ⚡ OpenAI Realtime API (Azure inclus)
- ⚡ Claude Messages
- ⚡ Google Gemini
- 🔄 Modèles Rerank (Cohere, Jina)

**Routage intelligent :**
- ⚖️ Aléatoire pondéré par canal
- 🔄 Nouvelle tentative automatique en cas d'échec
- 🚦 Limitation de débit par modèle et par utilisateur

---

## 🚢 Déploiement

### 📋 Prérequis

| Composant | Exigence |
|------|------|
| **Base de données** | SQLite (nœud unique) / MySQL ≥ 5.7.8 / PostgreSQL ≥ 9.6 |
| **Go** | 1.22+ |
| **Node.js** | 18+ (avec Bun) |

### ⚙️ Variables d'environnement

<details>
<summary>Variables courantes</summary>

| Variable | Description | Valeur par défaut |
|--------|------|--------|
| `SESSION_SECRET` | Clé de session (obligatoire pour multi-nœuds) | - |
| `CRYPTO_SECRET` | Clé de chiffrement (obligatoire avec Redis) | - |
| `SQL_DSN` | Chaîne de connexion à la base de données | - |
| `REDIS_CONN_STRING` | Chaîne de connexion Redis | - |
| `STREAMING_TIMEOUT` | Délai d'expiration du streaming (secondes) | `300` |
| `AZURE_DEFAULT_API_VERSION` | Version de l'API Azure | `2025-04-01-preview` |
| `ERROR_LOG_ENABLED` | Activation des logs d'erreur | `false` |

</details>

---

## 🔗 Projets liés

### Projets en amont

| Projet | Description |
|------|------|
| [new-api](https://github.com/QuantumNous/new-api) | Projet amont direct (AGPL v3) |
| [One API](https://github.com/songquanpeng/one-api) | Base du projet original |

---

## 💬 Aide et support

Toutes les contributions sont les bienvenues !

- 🐛 Signaler des bugs
- 💡 Proposer de nouvelles fonctionnalités
- 📝 Améliorer la documentation
- 🔧 Soumettre du code

---

## 📜 Licence

Ce projet est sous licence [GNU Affero General Public License v3.0 (AGPLv3)](./LICENSE).

Ce projet est un fork modifié de [new-api](https://github.com/QuantumNous/new-api) (AGPL v3), lui-même basé sur [One API](https://github.com/songquanpeng/one-api) (MIT License).

---

<div align="center">

### 💖 Merci d'utiliser 算和云 (SuanHeYun)

</div>
