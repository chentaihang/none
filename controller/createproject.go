package controller

import (
	"fmt"
	"math/rand"
	"my-go-project/internal/config"
	"my-go-project/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateProjectRequest struct {
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description"`
	StudentID      *int   `json:"student_id"`
	Status         string `json:"status" binding:"required"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	ProgressDate   string `json:"progress_date"`
	ProgressDesc   string `json:"progress_desc"`
	ProgressStatus string `json:"progress_status"`
	Type           string `json:"type"`
	Major          string `json:"major"`
}

// 创建项目处理函数
func CreateProject(c *gin.Context) {
	var request CreateProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误", "details": err.Error()})
		return
	}
	request.StudentID = nil
	request.ProgressStatus = "已提交"

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

	startDate, parseErr := time.Parse("2006-01-02", request.StartDate)
	if parseErr != nil {
		startDate = time.Now()
	}

	var endDate *time.Time
	if request.EndDate != "" {
		parsedEndDate, parseErr := time.Parse("2006-01-02", request.EndDate)
		if parseErr == nil {
			endDate = &parsedEndDate
		}
	}

	var progressDate *time.Time
	if request.ProgressDate != "" {
		parsedProgressDate, parseErr := time.Parse("2006-01-02", request.ProgressDate)
		if parseErr == nil {
			progressDate = &parsedProgressDate
		}
	}

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数
	randomInt := rand.Int()
	fmt.Println("随机整数:", randomInt)

	// 生成一个在 [0, 10000) 范围内的随机整数
	randomIntInRange := rand.Intn(10000)
	// projectID随机生成
	project := model.Project{
		ProjectID:      randomIntInRange,
		Title:          request.Title,
		Description:    request.Description,
		TeacherID:      teacher.TeacherID,
		StudentID:      request.StudentID,
		Status:         request.Status,
		StartDate:      startDate,
		EndDate:        endDate,
		ProgressDate:   progressDate,
		ProgressDesc:   request.ProgressDesc,
		ProgressStatus: &request.ProgressStatus,
		Type:           request.Type,
		Major:          request.Major,
	}
	// 如果课题存在
	if err := config.DB.Table("project").Where("title = ?", request.Title).First(&project).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "项目已存在"})
		return
	}
	if err := config.DB.Table("project").Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建项目失败"})
		return
	}

	c.JSON(http.StatusOK, project)
}
