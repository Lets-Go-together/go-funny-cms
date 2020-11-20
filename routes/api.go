package routes

import (
	"github.com/gin-gonic/gin"
	"gocms/app/http/index/controllers"
)

// 路由注册
func RegisterWebRoutes(router *gin.Engine) {

	// 定义 404 路由

	indexControler := new(controllers.IndexController)
	router.GET("/", indexControler.Index)
}
