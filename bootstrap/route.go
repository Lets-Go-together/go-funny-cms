package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gocms/app/http/middleware"
	"gocms/pkg/config"
	"gocms/routes"
)

// 初始化路由
func SteupRoute() {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(middleware.WebMiddleware)
	{
		routes.RegisterWebRoutes(router)
		routes.RegisterApiRoutes(router)
	}
	config.Router = router

	_ = router.Run(":" + config.GetString("APP_PORT"))
}
