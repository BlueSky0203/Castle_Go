package main

import (
	"log"
	"os"

	routes "Castle_Go/router"
	"Castle_Go/utils"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // ✅ 正確引入 swaggerFiles
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Castle_Go/docs" // ✅ 引入你的 Swagger 文檔
)

func main() {
	// 加載環境變數
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 嘗試連接資料庫
	utils.ConnectDatabase()

	// 創建 Gin 實例
	r := gin.Default()

	// 設定 Swagger UI 路由
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 設定你的其他路由
	routes.SetupRoutes(r)

	// 設置埠口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 啟動服務
	r.Run(":" + port)
}
