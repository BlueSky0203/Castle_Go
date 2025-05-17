package routes

import (
	handlers "Castle_Go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/castle-types", handlers.GetAllCastleTypes)
	r.GET("/castles", handlers.GetCastleListHandler)
	r.POST("/upload-castle-image", handlers.UploadCastleImage)
	r.POST("/create-castle", handlers.CreateCastle)
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
}
