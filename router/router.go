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

	//  JWT 中间件保护路由
	auth := r.Group("/")
	auth.Use(utils.JWT())
	{
		//  CreateProject 函数
		auth.POST("/project-create", controller.CreateProject)
		//  删除函数
		auth.DELETE("/delete-project", controller.DeleteProjectByTitle)
	}

	return r
}
