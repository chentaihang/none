package router

import (
	"my-go-project/controller"
	"my-go-project/pkg/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 登录处理函数
	r.POST("/login", controller.Login)

	// 注册处理函数
	r.POST("/register", controller.Register)

	// GetProjectByTeacher 函数
	r.GET("/project-by-name", controller.GetProjectByTeacher)
	// 获取所有的项目
	r.GET("/allproject", controller.GetAllProject)
	//  JWT 中间件保护路由
	auth := r.Group("/")
	auth.Use(utils.JWT())
	{
		//  CreateProject 函数
		auth.POST("/project-create", controller.CreateProject)
		//  删除函数
		auth.DELETE("/delete-project", controller.DeleteProjectByTitle)
		//  更新函数
		auth.PUT("/update-project", controller.UpdateProject)
		// 查询已提交的项目
		auth.Handle("GET", "/updatestatus", controller.QuerySubmittedProjects)
		// 更新项目状态
		auth.Handle("PUT", "/updatestatus", controller.UpdateProjectStatus)
		// 获取个人信息(token匹配使用)
		auth.GET("/personinfo", controller.GetPersonInfo)
		// 获取所有学生信息
		auth.GET("/allstudentinfo", controller.GetALLstudentinfo)
	}

	return r
}
