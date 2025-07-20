# 🏰 Castle Go API

## 📘 專案簡介

這是一個使用 **Go** 和 **Gin** 框架開發的後端 API 專案，支援圖片上傳、WebSocket 即時通訊功能，以及使用 Swagger 文件進行介面測試。此專案可用於圖像管理、儀表板即時監控、通知推播等場景。

---

## 🚀 功能列表

- ✅ 圖片上傳至 Cloudinary
- ✅ WebSocket 即時連線與訊息廣播
- ✅ Swagger UI 文件與 API 測試介面
- ✅ 使用 .env 設定環境變數（資料庫、Cloudinary 等）

---

## 🧱 技術棧

| 類別         | 技術                          |
|--------------|-------------------------------|
| 語言         | Go 1.18+                       |
| Web 框架     | [Gin](https://gin-gonic.com/) |
| ORM          | [GORM](https://gorm.io/)      |
| 圖片儲存     | [Cloudinary](https://cloudinary.com/) |
| 文件生成     | [Swagger + swaggo](https://github.com/swaggo/gin-swagger) |
| 即時通訊     | WebSocket (原生 + gorilla/websocket) |
| 環境設定     | [godotenv](https://github.com/joho/godotenv) |

---

## ⚙️ 安裝與啟動
- http://localhost:8080/docs/index.html

### 📦 前置需求

- Go 1.18 或以上
- 已註冊並設定 Cloudinary 帳戶
- 本地安裝 PostgreSQL 或其他支援的資料庫（若使用）

### 🛠️ 設定 `.env`

```env
PORT=8080
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_secret
