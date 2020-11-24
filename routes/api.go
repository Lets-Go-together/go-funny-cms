package routes

import (
	"github.com/gin-gonic/gin"
	adminContro "gocms/app/http/admin/controllers"
	"gocms/app/http/index/controllers"
)

// 路由注册
func RegisterWebRoutes(router *gin.Engine) {

	// 定义 404 路由

	indexControler := new(controllers.IndexController)
	router.GET("/", indexControler.Index)

	authController := new(adminContro.AuthController)
	router.POST("/admin/login", authController.Login)
}
