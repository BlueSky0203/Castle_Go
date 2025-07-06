package models

import (
	"Castle_Go/utils"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

// GORM 模型：自動處理 ID、CreatedAt、UpdatedAt 等欄位
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Name         string    `json:"name"`
}

// 這段讓 GORM 使用 "user" 而不是預設的 "users"
func (User) TableName() string {
	return "user"
}

func AuthenticateUser(db *gorm.DB, identifier, password string) (*User, error) {
	var user User

	// 嘗試用 username 或 email 查詢
	if err := db.Where("username = ? OR email = ?", identifier, identifier).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// 驗證密碼
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

// 註冊
func CreateUser(db *gorm.DB, username, email, password, name string) (*User, error) {
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func AuthenticateFirebase(db *gorm.DB, email, name string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err == nil {
		// 找到就回傳
		return &user, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 其他錯誤回傳
		return nil, err
	}

	// 沒找到，新增 user
	username := strings.Split(email, "@")[0]
	user = User{
		Username:     username,
		Email:        email,
		PasswordHash: "", // Firebase 登入不用密碼
		Name:         name,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
