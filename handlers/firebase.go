package handlers

import (
	"Castle_Go/models"
	"Castle_Go/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FirebaseLoginInput struct {
	FirebaseToken string `json:"firebase_token" binding:"required"`
}

func FirebaseLogin(c *gin.Context) {
	var input FirebaseLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Firebase token is required"})
		return
	}

	// 驗證 Firebase Token
	token, err := utils.FirebaseAuth.VerifyIDToken(context.Background(), input.FirebaseToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Firebase token"})
		return
	}

	email, _ := token.Claims["email"].(string)
	name, _ := token.Claims["name"].(string)
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found in Firebase token"})
		return
	}

	// 查詢或創建使用者
	user, err := models.AuthenticateFirebase(utils.DB, email, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// 產生你自己的 JWT
	jwtToken, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
		"user":  user,
	})
}
