package routes

import (
	"github.com/gin-gonic/gin"
	"gocms/app/http/index/controllers"
	"gocms/app/http/middleware"
)

func RegisterWebRoutes(router *gin.Engine) {
	// 前台项目
	webRouter := router.Group("/")

	webRouter.Use(middleware.DefaultMiddle)
	{
		indexControler := new(controllers.IndexController)
		router.GET("/", indexControler.Index)
	}
}
