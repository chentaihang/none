package controller

import (
	"net/http"

	"my-go-project/internal/config"
	"my-go-project/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 返回所有项目列表
func GetAllProject(c *gin.Context) {
	var project []model.Project
	result := config.DB.Table("project").Find(&project)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "项目未找到"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询项目失败" + result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, project)
}
