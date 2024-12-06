package controller

import (
	"net/http"

	"my-go-project/internal/config"
	"my-go-project/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 查询已提交的项目
func QuerySubmittedProjects(c *gin.Context) {
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权的访问"})
		return
	}
	var temp string = "已提交"
	var projects []model.Project
	if err := config.DB.Table("project").Where("progress_status = ?", temp).Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询项目失败"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// 更新已提交的项目为已审核
func UpdateProjectStatus(c *gin.Context) {
	ProjectID := c.Query("project_id")
	if ProjectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "项目ID不能为空"})
		return
	}
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权的访问"})
		return
	}
	// 检查项目是否存在
	var project model.Project
	if err := config.DB.Table("project").Where("project_id = ?", ProjectID).First(&project).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "所要更新的项目未找到"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询项目失败" + err.Error()})
		}
		return
	}
	var temp string = "已审核"
	project.ProgressStatus = &temp
	if err := config.DB.Table("project").Where("project_id = ?", ProjectID).Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新项目失败"})
		return
	}

	c.JSON(http.StatusOK, project)
}
