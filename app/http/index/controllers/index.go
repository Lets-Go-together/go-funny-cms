package controllers

import (
	"gocms/pkg/auth"
	"gocms/wrap"
)

type IndexController struct{}

var (
	jwtAction = auth.JwtAction{}
)

func (*IndexController) Index(c *wrap.ContextWrapper) {
	//token := jwtAction.GetToken(users.AuthUser{
	//	Id:   10,
	//	Name: "chenf",
	//})
	//user, err := jwtAction.ParseToken(token)
	//if err != nil {
	//	c.JSON(200, gin.H{
	//		"code": 403,
	//		"msg":  "解析失败: " + err.Error(),
	//	})
	//	return
	//}
	//
	//c.JSON(200, gin.H{
	//	"code": 200,
	//	"user": user,
	//})
}
