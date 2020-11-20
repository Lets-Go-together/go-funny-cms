package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gocms/routes"
)

// 初始化路由
func SteupRoute() {
	r := gin.Default()

	routes.RegisterWebRoutes(r)
}
