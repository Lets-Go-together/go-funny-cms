package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gocms/app/models/admin"
	"gocms/pkg/auth"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/logger"
	"gocms/pkg/response"
)

var jwtAction = auth.JwtAction{}

func Authenticate(c *gin.Context) {
	token := c.GetHeader("authorization")
	token = auth.ValidateToken(token)

	if len(token) == 0 {
		response.ErrorResponse(401, "鉴权失败").WriteTo(c)
		c.Abort()
	}

	adminUser, err := jwtAction.ParseToken(token)
	if err != nil {
		logger.Info("Token 已失效", err.Error())
		response.ErrorResponse(401, "Token 已失效").WriteTo(c)
		c.Abort()
	}

	adminUser.Roles = admin.GetRoles(adminUser.Account)
	adminUser.Menus = admin.GetMenus(adminUser.Roles, adminUser.Account)
	adminUser.Permissions = admin.GetPermissions(adminUser.Account)
	admin.AuthUser = adminUser

	fmt.Println(adminUser)
	// 权限检查
	permission := c.FullPath()
	method := c.Request.Method
	if !rabc.AllowPermission(adminUser.Account, permission, method) {
		response.ErrorResponse(403, "您没有权限操作").WriteTo(c)
		c.Abort()
	}

	c.Next()
}
