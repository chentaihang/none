package controller

import (
	"net/http"

	"my-go-project/internal/config"
	"my-go-project/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取个人信息
func GetPersonInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权的访问"})
		return
	}

	var user model.User
	if err := config.DB.Table("user").Where("user_id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户信息未找到"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询用户信息失败", "details": err.Error()})
		}
		return
	}

	if user.UserType == "student" {
		var student model.Student
		if err := config.DB.Table("student").Where("student_id = ?", userID).First(&student).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "学生信息未找到"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "查询学生信息失败", "details": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, student)
	} else if user.UserType == "teacher" {
		var teacher model.Teacher
		if err := config.DB.Table("teacher").Where("teacher_id = ?", userID).First(&teacher).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "教师信息未找到"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "查询教师信息失败", "details": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, teacher)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "未知的用户类型"})
	}
}

type ProjectStudent struct {
	ProjectID int `json:"project_id"`
	StudentID int `json:"student_id"`
}

type ReturnInfo struct {
	Student model.Student `json:"student"`
	Status  string        `json:"status"`
}

// 获取所有学生信息
func GetALLstudentinfo(c *gin.Context) {
	var students []model.Student
	var projectStudents []ProjectStudent
	var returnInfos []ReturnInfo

	config.DB.Table("project_student").Find(&projectStudents)

	for _, ps := range projectStudents {
		var student model.Student
		if err := config.DB.Table("student").Where("student_id = ?", ps.StudentID).First(&student).Error; err != nil {
			continue
		}
		returnInfos = append(returnInfos, ReturnInfo{
			Student: student,
			Status:  "已选题",
		})
	}

	// 查找未选题的学生
	config.DB.Table("student").Not("student_id IN (?)", config.DB.Table("project_student").Select("student_id")).Find(&students)
	for _, student := range students {
		returnInfos = append(returnInfos, ReturnInfo{
			Student: student,
			Status:  "未选题",
		})
	}

	c.JSON(http.StatusOK, returnInfos)
}
