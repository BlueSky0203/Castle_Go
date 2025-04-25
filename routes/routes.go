package routes

import (
	"Castle_Go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/upload", handlers.UploadCastleImage)
}
