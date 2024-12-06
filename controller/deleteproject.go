package controller

import (
	"net/http"

	"my-go-project/internal/config"
	"my-go-project/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 删除教师对应的项目（通过标题来确定）
func DeleteProjectByTitle(c *gin.Context) {
	title := c.Query("title")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 title 参数"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权的访问"})
		return
	}

	var teacher model.Teacher
	if err := config.DB.Table("teacher").Where("user_id = ?", userID).First(&teacher).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "教师未找到"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询教师失败"})
		}
		return
	}

	result := config.DB.Table("project").Where("teacher_id = ? AND title = ?", teacher.TeacherID, title).Delete(&model.Project{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除项目失败+" + result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "项目未找到"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "项目删除成功"})
}
