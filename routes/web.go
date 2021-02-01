package routes

import (
	"github.com/gin-gonic/gin"
	Index "gocms/app/http/index/controllers"
	"gocms/app/http/middleware"
	"gocms/wrap"
)

func RegisterWebRoutes(router *gin.Engine) {
	// 前台项目
	webRouter := router.Group("/")

	webRouter.Use(middleware.DefaultMiddle)
	{
		IndexController := new(Index.IndexController)
		router.GET("/", func(context *gin.Context) {
			IndexController.Index(wrap.Context(context))
		})
	}
}
