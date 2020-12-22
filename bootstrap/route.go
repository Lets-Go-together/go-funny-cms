package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gocms/app/http/middleware"
	"gocms/pkg/config"
	"gocms/routes"
)

// 初始化路由
func SteupRoute(params ...interface{}) {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(middleware.WebMiddleware)
	{
		routes.RegisterWebRoutes(router)
		routes.RegisterApiRoutes(router)
	}
	config.Router = router

	if len(params) > 0 {
		if cast.ToBool(params[0]) == true {
			return
		}
	}
	_ = router.Run(":" + config.GetString("APP_PORT"))
}
