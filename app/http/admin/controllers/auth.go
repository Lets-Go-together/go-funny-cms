package controllers

import (
	"github.com/gin-gonic/gin"
	"gocms/app/models/admin"
	authValidate "gocms/app/validates/auth"
	"gocms/pkg/auth"
	"gocms/pkg/config"
	"gocms/pkg/response"
)

type AuthController struct{}

var (
	jwtAction = auth.JwtAction{}
)

func (*AuthController) Login(c *gin.Context) {
	authValidateAction := authValidate.LoginAction{}
	if msg := authValidateAction.Validate(c); len(msg) > 0 {
		response.ErrorResponse(403, msg).WriteTo(c)
		return
	}

	params := authValidateAction.GetLoginData()
	adminModel := admin.Admin{}
	config.Db.Where("account = ?", params.Account).Find(&adminModel)

	if adminModel.ID == 0 {
		response.ErrorResponse(403, "用户不存在").WriteTo(c)
		return
	}

	password := adminModel.Password
	if !auth.ValidatePassword(password, params.Password) {
		response.ErrorResponse(403, "密码错误").WriteTo(c)
		return
	}

	authAdmin := admin.GetAuthAdmin(adminModel)
	token := jwtAction.GetToken(authAdmin)

	response.SuccessResponse(map[string]string{
		"token": token,
	}).WriteTo(c)
}

// 我的信息
func (*AuthController) Me(c *gin.Context) {
	user := config.AuthAdmin
	response.SuccessResponse(user).WriteTo(c)

	return
}

// 注销
func (*AuthController) Logout(c *gin.Context) {
	user := config.AuthAdmin
	response.SuccessResponse(user).WriteTo(c)
}

func (*AuthController) Register(c *gin.Context) {

}
