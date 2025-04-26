package model

import "time"

type Castle struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	NameZh      string    `json:"name_zh"`
	Description string    `json:"description"`
	Country     string    `json:"country"`
	Type        int       `json:"type"`
	ImageURL    string    `json:"image_url"`
	BuildYear   int       `json:"build_year"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Castle) TableName() string {
	return "castle"
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
