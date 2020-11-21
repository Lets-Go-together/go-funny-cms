package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gocms/pkg/config"
	"gocms/routes"
)

// 初始化路由
func SteupRoute() {
	router := gin.Default()
	routes.RegisterWebRoutes(router)

	_ = router.Run(":" + config.GetEnv("APP_PORT"))
}
