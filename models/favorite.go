package models

import (
	"Castle_Go/utils"
	"time"
)

type Favorite struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id" gorm:"uniqueIndex:idx_user_castle"`
	CastleId  uint      `json:"castle_id" gorm:"uniqueIndex:idx_user_castle"`
	Type      int       `json:"type" gorm:"default:1"`                           // 1: 收藏中, 0: 取消收藏
	Castle    Castle    `json:"castle" gorm:"foreignKey:CastleId;references:ID"` // 關聯
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func GetFavoriteList(userID uint) ([]Favorite, error) {
	var favorite []Favorite
	err := utils.DB.
		Where("user_id = ?", userID).
		Preload("Castle").
		Find(&favorite).Error
	return favorite, err
}

func (Favorite) TableName() string {
	return "favorite"
}
