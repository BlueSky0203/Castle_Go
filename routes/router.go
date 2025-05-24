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
		auth.GET("/favorites", handlers.GetFavorites)
		auth.POST("/favorites", handlers.AddFavorite)
		auth.DELETE("/favorites/:castle_id", handlers.RemoveFavorite)
	}
}
