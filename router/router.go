package routes

import (
	handlers "Castle_Go/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/upload", handlers.UploadCastleImage)
	r.GET("/castle-types", handlers.GetAllCastleTypes)
}
