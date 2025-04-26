package handlers

import (
	"fmt"
	"net/http"

	"Castle_Go/model"
	"Castle_Go/utils"

	"github.com/gin-gonic/gin"
)

// UploadCastleImage 上傳城堡圖片
// @Summary 上傳城堡圖片
// @Description 上傳一張圖片並上傳至 Cloudinary，回傳圖片 URL
// @Tags 圖片上傳
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "要上傳的圖片"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /upload [post]
func UploadCastleImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image found"})
		return
	}

	// 存到本地暫存
	filePath := fmt.Sprintf("./%s", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload failed"})
		return
	}

	// 上傳到 Cloudinary
	uploadResult, err := utils.UploadImage(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary upload failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Upload successful",
		"imageUrl": uploadResult.SecureURL,
	})
}

// GetAllCastleTypes 取得所有城堡類型
// @Summary 取得所有城堡類型
// @Tags CastleType
// @Produce json
// @Success 200 {array} model.CastleType
// @Router /castle-types [get]
func GetAllCastleTypes(c *gin.Context) {
	var types []model.CastleType

	if err := utils.DB.Find(&types).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法取得城堡類型資料"})
		return
	}
	c.JSON(http.StatusOK, types)
}
