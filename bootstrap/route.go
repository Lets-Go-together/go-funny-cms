package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gocms/pkg/config"
	"gocms/routes"
)

// 初始化路由
func SteupRoute() {
	router := gin.New()
	router.Use(gin.Logger())

	routes.RegisterWebRoutes(router)
	routes.RegisterApiRoutes(router)

	_ = router.Run(":" + config.GetString("APP_PORT"))
}
