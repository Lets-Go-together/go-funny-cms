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
	data := gin.H{
		"msg":   "Hello Word",
		"token": token,
		"parse": jwtAction.ParseToken(token),
	}
	c.JSON(403, data)
}
