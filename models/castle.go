package models

import (
	"Castle_Go/utils"
	"time"
)

type Castle struct {
	ID          uint      `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Name        string    `json:"name"`
	NameZh      string    `json:"name_zh"`
	Description string    `json:"description"`
	Country     string    `json:"country"`
	Type        int       `json:"type"`
	ImageURL    string    `json:"image_url"`
	BuildYear   int       `json:"build_year"`
	CreatedAt   time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"-" gorm:"autoUpdateTime"`
}

func (Castle) TableName() string {
	return "castle"
}

func GetCastleList(page, pageSize int) ([]Castle, int64, error) {
	var castles []Castle
	var total int64

	// 查總數
	if err := utils.DB.Model(&Castle{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分頁查資料
	offset := (page - 1) * pageSize
	err := utils.DB.
		Limit(pageSize).
		Offset(offset).
		Find(&castles).Error

	return castles, total, err
}

type CastleType struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	NameZh      string    `json:"name_zh"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (CastleType) TableName() string {
	return "castle_type"
}
