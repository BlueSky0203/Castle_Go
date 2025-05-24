package handlers

import (
	"Castle_Go/models"
	"Castle_Go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetFavoriteList 取得我的收藏列表
// @Summary 取得收藏列表
// @Description 取得目前使用者的所有收藏城堡
// @Tags 收藏
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /favorites [get]
// @Security BearerAuth
func GetFavoriteList(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	favorites, err := models.GetFavoriteList(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch favorites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"favorites": favorites})
}

type AddFavoriteInput struct {
	CastleId uint `json:"castle_id" binding:"required"`
}

// AddFavorite 新增收藏
// @Summary 新增收藏
// @Description 新增或恢復一筆使用者的收藏城堡資料，若已存在則恢復為收藏狀態
// @Tags 收藏
// @Accept json
// @Produce json
// @Param favorite body AddFavoriteInput true "收藏資料"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /favorites [post]
// @Security BearerAuth
func AddFavorite(c *gin.Context) {
	var input AddFavoriteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.MustGet("userID").(uint)

	var fav models.Favorite
	err := utils.DB.Where("user_id = ? AND castle_id = ?", userID, input.CastleId).
		First(&fav).Error

	if err == nil {
		// 已存在紀錄，更新為 type=1（恢復收藏）
		if fav.Type == 1 {
			c.JSON(http.StatusOK, gin.H{"message": "Already favorited"})
			return
		}
		if err := utils.DB.Model(&fav).Update("type", 1).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore favorite"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Favorite restored"})
		return
	}

	// 不存在紀錄，新增
	newFav := models.Favorite{
		UserId:   userID,
		CastleId: input.CastleId,
		Type:     1,
	}
	if err := utils.DB.Create(&newFav).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add favorite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorite added"})
}

// RemoveFavorite 取消收藏
// @Summary 取消收藏
// @Description 將使用者對指定城堡的收藏標記為取消（type 設為 0）
// @Tags 收藏
// @Produce json
// @Param castle_id path int true "城堡 ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /favorites/{castle_id} [delete]
// @Security BearerAuth
func RemoveFavorite(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	castleIDStr := c.Param("castle_id")

	castleID, err := strconv.Atoi(castleIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid castle ID"})
		return
	}

	// 更新為 type = 0，而不是刪除
	if err := utils.DB.Model(&models.Favorite{}).
		Where("user_id = ? AND castle_id = ?", userID, castleID).
		Update("type", 0).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove favorite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorite removed"})
}
