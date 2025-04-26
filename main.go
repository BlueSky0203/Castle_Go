package main

import (
	"log"
	"os"
	"time"

	"Castle_Go/routes"
	"Castle_Go/utils"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Castle_Go/docs"
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

	// 加上 CORS 中介軟體
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 允許前端開發機器存取
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
