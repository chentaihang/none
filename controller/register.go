package controller

import (
	"net/http"

	"my-go-project/internal/config"
	"my-go-project/internal/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 注册处理函数
func Register(c *gin.Context) {
	var usertmp model.User
	if err := c.ShouldBindJSON(&usertmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	// 检查用户是否已经存在
	var existingUser model.User
	if err := config.DB.Table("user").Where("user_id = ? OR username = ?", usertmp.UserID, usertmp.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "用户已存在"})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usertmp.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}

	user := model.User{
		UserID:   usertmp.UserID,
		Username: usertmp.Username,
		Password: string(hashedPassword),
		Role:     usertmp.Role,
		Email:    usertmp.Email,
		UserType: usertmp.UserType,
	}

	if err := config.DB.Table("user").Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户注册失败"})
		return
	}

	// 注册成功，返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"user_id":   user.UserID,
		"username":  user.Username,
		"role":      user.Role,
		"email":     user.Email,
		"user_type": user.UserType,
	})
}
