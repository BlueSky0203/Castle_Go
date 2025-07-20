package main

import (
	"Castle_Go/websocket"
	"fmt"
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

// @title Castle API
// @version 1.0
// @description Castle API for managing castles.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// 加載 .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing without it.")
	}

	if err := utils.InitRedis(); err != nil {
		log.Fatalf("Failed to init Redis: %v", err)
	}

	// 初始化 Firebase & DB
	utils.InitFirebase()
	utils.ConnectDatabase()

	// Gin 實例
	r := gin.Default()

	// CORS
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://castle-99e47.web.app", "https://castle-99e47.firebaseapp.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	r.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/ws" {
			c.Next()
			return
		}
		corsMiddleware(c)
	})

	// WebSocket Hub
	hub := websocket.NewHub()
	go hub.Run()

	// WebSocket 加入到 Gin 路由
	r.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(hub, c.Writer, c.Request)
	})

	// Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 其他 API 路由
	routes.SetupRoutes(r)

	// 啟動伺服器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("🚀 Server listening on port %s\n", port)
	r.Run(":" + port)
}
