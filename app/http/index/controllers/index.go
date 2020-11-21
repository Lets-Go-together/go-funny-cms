package controllers

import (
	"github.com/gin-gonic/gin"
	"gocms/app/models/users"
	"gocms/pkg/auth"
)

type IndexController struct{}

var (
	jwtAction = auth.JwtAction{}
)

func (*IndexController) Index(c *gin.Context) {
	token := jwtAction.GetToken(users.AuthUser{
		Id:   10,
		Name: "chenf",
	})
	user, err := jwtAction.ParseToken(token)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "解析失败: " + err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"user": user,
	})
}
