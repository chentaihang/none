package controller

import (
	"net/http"

	"my-go-project/internal/config"
	"my-go-project/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 查找教师对应的项目（通过标题查找）
func GetProjectByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 title 参数"})
		return
	}

	var project model.Project
	result := config.DB.Table("project").Where("title = ?", title).First(&project)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "项目未找到"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询项目失败"})
		}
		return
	}

	c.JSON(http.StatusOK, project)
}
