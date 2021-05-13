package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gocms/app/models/admin"
	"gocms/app/models/menu"
	"gocms/pkg/auth"
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

	menus := admin.GetMenus(adminUser.Roles, adminUser.Account)
	fmt.Println(len(menus))
	adminUser.Menus = make([]menu.MenuRouter, len(menus))
	for k, menu := range menus {
		adminUser.Menus[k] = menu
	}

	adminUser.Roles = admin.GetRoles(adminUser.Account)
	adminUser.Permissions = admin.GetPermissions(adminUser.Account)
	admin.AuthUser = adminUser

	c.Next()
}
