package controller

import (
	"my-go-project/internal/config"
	"my-go-project/internal/model"
	"my-go-project/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	UserID   string `json:"userID" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	var user model.User
	if err := config.DB.Table("user").Where("user_id = ?", loginRequest.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 验证密码
	// 判断密码是否加密：如果密码长度大于 20，则说明密码已经加密
	if len(user.Password) < 20 {

		if user.Password != loginRequest.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
			return
		}
	} else {
		// 验证加密密码
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
			return
		}
	}
	// 生成 JWT
	token, err := utils.GenerateToken(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 Token 失败"})
		return
	}
	// 登录成功，返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"user_id":   user.UserID,
		"username":  user.Username,
		"role":      user.Role,
		"email":     user.Email,
		"user_type": user.UserType,
		"token":     token,
	})
}
