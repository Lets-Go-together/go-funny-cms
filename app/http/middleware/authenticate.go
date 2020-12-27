package middleware

import (
	"github.com/gin-gonic/gin"
	"gocms/app/models/admin"
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

	user, err := jwtAction.ParseToken(token)
	if err != nil {
		logger.Info("Token 已失效", err.Error())
		response.ErrorResponse(401, "Token 已失效").WriteTo(c)
		c.Abort()
	}

	admin.AuthUser = user
	c.Next()
}
