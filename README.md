# 🏰 Castle Go API

## 📘 專案簡介

這是一個使用 **Go** 和 **Gin** 框架開發的後端 API 專案，支援圖片上傳、WebSocket 即時通訊功能，並整合 Upstash Redis 進行訊息廣播。此專案可用於圖像管理、儀表板即時監控、通知推播等場景，搭配前端 Vue SPA 可快速部署完整應用。

---

## 🚀 功能列表

- ✅ 圖片上傳至 Cloudinary
- ✅ WebSocket 即時連線與訊息廣播  
  - 使用 gorilla/websocket 架構  
  - 整合 Upstash Redis Pub/Sub，支援多源訊息分發給所有 WebSocket 客戶端
- ✅ Swagger UI 文件與 API 測試介面
- ✅ 使用 `.env` 設定環境變數（資料庫、Cloudinary、Redis 等）

---

## 🧱 技術棧

| 類別         | 技術                                                         |
|--------------|--------------------------------------------------------------|
| 語言         | Go 1.18+                                                     |
| Web 框架     | [Gin](https://gin-gonic.com/)                                |
| ORM          | [GORM](https://gorm.io/)                                     |
| 圖片儲存     | [Cloudinary](https://cloudinary.com/)                        |
| 文件生成     | [Swagger + swaggo](https://github.com/swaggo/gin-swagger)    |
| 即時通訊     | WebSocket (gorilla/websocket) + Redis Pub/Sub (Upstash)      |
| 訊息代理     | [Upstash Redis](https://upstash.com/)                        |
| 環境設定     | [godotenv](https://github.com/joho/godotenv)                 |

---

## ⚙️ 安裝與啟動

### 📦 前置需求

- Go 1.18 或以上版本
- PostgreSQL（或你設計的其他資料庫）
- 已註冊並設定 Cloudinary 帳戶
- Upstash Redis 帳戶

---

### 🛠️ `.env` 環境變數設定範例

請建立 `.env` 檔案並填入以下設定：

```env
# ✅ 伺服器設定
PORT=8080
JWT_SECRET=<your_jwt_secret>

# ✅ Supabase 資料庫
SUPABASE_DB_URL=postgresql://<user>:<password>@<host>:<port>/<database>

# ✅ Firebase（前端用）
VITE_FIREBASE_API_KEY=<your_firebase_api_key>
VITE_FIREBASE_AUTH_DOMAIN=<your_project>.firebaseapp.com
VITE_FIREBASE_PROJECT_ID=<your_project_id>

# ✅ Cloudinary 圖片上傳
CLOUDINARY_URL=cloudinary://<api_key>:<api_secret>@<cloud_name>
CLOUDINARY_API_KEY=<your_api_key>
CLOUDINARY_API_SECRET=<your_api_secret>
CLOUDINARY_CLOUD_NAME=<your_cloud_name>

# ✅ Upstash Redis
UPSTASH_REDIS_URL=rediss://default:<your_token>@<your_instance>.upstash.io:6379
