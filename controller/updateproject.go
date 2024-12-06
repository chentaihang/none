package controller

import (
	"net/http"
	"time"

	"my-go-project/internal/config"
	"my-go-project/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 更新项目请求结构体
type UpdateProjectRequest struct {
	ProjectID      int     `json:"project_id" binding:"required"`
	Title          string  `json:"title"`
	Description    *string `json:"description"`
	StudentID      *int    `json:"student_id"`
	Status         *string `json:"status"`
	StartDate      *string `json:"start_date"`
	EndDate        *string `json:"end_date"`
	ProgressDate   *string `json:"progress_date"`
	ProgressDesc   *string `json:"progress_desc"`
	ProgressStatus *string `json:"progress_status"`
}

// 更新项目处理函数
func UpdateProject(c *gin.Context) {
	var request UpdateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误", "details": err.Error()})
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

	var project model.Project
	if err := config.DB.Table("project").Where("project_id = ?", request.ProjectID).First(&project).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "所要更新的项目未找到"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询项目失败"})
		}
		return
	}

	// 更新项目字段
	if request.Title != "" {
		project.Title = request.Title
	}
	if request.Description != nil {
		project.Description = *request.Description
	}
	if request.StudentID != nil {
		project.StudentID = request.StudentID
	}
	if request.Status != nil {
		project.Status = *request.Status
	}
	if request.StartDate != nil {
		startDate, err := time.Parse("2006-01-02", *request.StartDate)
		if err == nil {
			project.StartDate = startDate
		}
	}
	if request.EndDate != nil {
		endDate, err := time.Parse("2006-01-02", *request.EndDate)
		if err == nil {
			project.EndDate = &endDate
		}
	}
	if request.ProgressDate != nil {
		progressDate, err := time.Parse("2006-01-02", *request.ProgressDate)
		if err == nil {
			project.ProgressDate = &progressDate
		}
	}
	if request.ProgressDesc != nil {
		project.ProgressDesc = *request.ProgressDesc
	}
	if err := config.DB.Table("project").Where("project_id = ? AND teacher_id = ? ", project.ProjectID, teacher.TeacherID).Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新项目失败" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}
