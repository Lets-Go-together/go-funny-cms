package controllers

import (
	"github.com/gin-gonic/gin"
	"gocms/pkg/auth"
)

type IndexController struct{}

var (
	jwtAction = auth.JwtAction{}
)

func (*IndexController) Index(c *gin.Context) {
	token := jwtAction.GetToken()
	data := gin.H{
		"msg":   "Hello Word",
		"token": token,
		"parse": jwtAction.ParseToken(token),
	}
	c.JSON(403, data)
}
