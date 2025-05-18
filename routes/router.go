package routes

import (
	"Castle_Go/handlers"
	"Castle_Go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 不需要登入就能使用的
	r.GET("/castle-types", handlers.GetAllCastleTypes)
	r.GET("/castles", handlers.GetCastleListHandler)
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	// 需要登入後才可用的
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.POST("/upload-castle-image", handlers.UploadCastleImage)
		auth.POST("/create-castle", handlers.CreateCastle)
		// 其他需要驗證的 CRUD 都可以加在這
	}
}
