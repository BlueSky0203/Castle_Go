package model

import "time"

type Favorite struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id" gorm:"primaryKey"`
	CastleId  uint      `json:"castle_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (Favorite) TableName() string {
	return "favorite"
}
