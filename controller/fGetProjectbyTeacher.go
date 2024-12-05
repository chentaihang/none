package controller

import (
	"net/http"

	"my-go-project/internal/config"
	"my-go-project/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 查找教师对应的项目（通过教师名称查找）
func GetProjectByTeacher(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 name 参数"})
		return
	}
	var teacher model.Teacher
	err := config.DB.Table("teacher").Where("name = ?", name).First(&teacher)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "教师未找到"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询教师失败"})
		}
		return
	}

	var project []model.Project
	result := config.DB.Table("project").Where("teacher_id = ?", teacher.TeacherID).Find(&project)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询项目失败"})
		return
	}

	if len(project) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "项目未找到"})
		return
	}

	c.JSON(http.StatusOK, project)
}
