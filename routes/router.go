package routes

import (
	handlers "Castle_Go/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/upload-castle-image", handlers.UploadCastleImage)
	r.POST("/create-castle", handlers.CreateCastle)
	r.GET("/castle-types", handlers.GetAllCastleTypes)
}
