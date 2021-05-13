package middleware

import (
	"github.com/gin-gonic/gin"
	"gocms/app/models/admin"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/response"
)

func Permission(c *gin.Context) {
	adminUser := admin.AuthUser
	// 权限检查
	permission := c.FullPath()
	method := c.Request.Method
	if !rabc.AllowPermission(adminUser.Account, permission, method) {
		response.ErrorResponse(403, "您没有权限操作").WriteTo(c)
		c.Abort()
	}

	c.Next()
}
