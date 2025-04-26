package handlers

import (
	"fmt"
	"net/http"

	models "Castle_Go/models"
	"Castle_Go/utils"

	"github.com/gin-gonic/gin"
)

// UploadCastleImage 上傳圖片並返回圖片 URL
// @Summary 上傳城堡圖片
// @Description 上傳城堡圖片並儲存至 Cloudinary，返回圖片 URL
// @Tags 圖片上傳
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "城堡圖片"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /upload-castle-image [post]
func UploadCastleImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image found"})
		return
	}

	// 存到本地暫存
	filePath := fmt.Sprintf("./%s", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Image save failed"})
		return
	}

	// 上傳到 Cloudinary
	uploadResult, err := utils.UploadImage(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary upload failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": uploadResult.SecureURL})
}

// CreateCastle 新增一筆城堡資料
// @Summary 新增城堡
// @Description 新增一筆城堡資料
// @Tags 城堡
// @Accept json
// @Produce json
// @Param castle body models.Castle true "城堡資料"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /create-castle [post]
func CreateCastle(c *gin.Context) {
	var castle models.Castle

	// 解析 JSON
	if err := c.ShouldBindJSON(&castle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 存進資料庫
	if err := utils.DB.Create(&castle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create castle"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Castle created successfully", "castle": castle})
}

// GetAllCastleTypes 取得所有城堡類型
// @Summary 取得所有城堡類型
// @Tags CastleType
// @Produce json
// @Success 200 {array} model.CastleType
// @Router /castle-types [get]
func GetAllCastleTypes(c *gin.Context) {
	var types []models.CastleType

	if err := utils.DB.Find(&types).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve castle types"})
		return
	}
	c.JSON(http.StatusOK, types)
}
