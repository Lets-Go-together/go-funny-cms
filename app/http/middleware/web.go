package middleware

import (
	"github.com/gin-gonic/gin"
	"gocms/pkg/config"
)

func WebMiddleware(c *gin.Context) {
	// 注入request 到全局
	config.Request = c.Request
	c.Next()
}
