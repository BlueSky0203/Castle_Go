package handlers

import (
	models "Castle_Go/models"
	"Castle_Go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name"`
}

// Register 使用者註冊
// @Summary 使用者註冊
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body RegisterInput true "註冊資料"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.CreateUser(utils.DB, input.Username, input.Email, input.Password, input.FullName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "registration successful",
		"user":    user,
	})
}

type LoginInput struct {
	Identifier string `json:"identifier" binding:"required"` // 可以是 username 或 email
	Password   string `json:"password" binding:"required"`
}

// Login 使用者登入
// @Summary 使用者登入
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body LoginInput true "登入資料"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.AuthenticateUser(utils.DB, input.Identifier, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
