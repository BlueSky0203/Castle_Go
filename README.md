# Castle Go API

## 專案簡介

這是一個使用 Go 和 Gin 框架構建的後端 API 服務。專案提供圖片上傳至 Cloudinary 的功能，並支持查看和管理相關資源。此專案可用於圖像管理、文件上傳等用途。

## 功能

- **圖片上傳**：提供 API 接口，將圖片上傳至 Cloudinary。
- **Swagger UI**：內建 Swagger 介面，方便進行 API 測試。
- **環境變數配置**：使用 `.env` 檔案配置 API 密鑰、Cloudinary 資訊等環境變數。

## 技術棧

- **Go**：後端語言
- **Gin**：Web 框架
- **Cloudinary**：圖片上傳及管理
- **Swagger**：API 文檔生成

## 安裝與設置

### 前置需求

- Go 1.18 或更高版本
- 設定 Cloudinary 帳戶並取得 API 金鑰