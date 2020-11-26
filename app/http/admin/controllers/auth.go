package controllers

import (
	"github.com/gin-gonic/gin"
	"gocms/app/models/admin"
	authValidate "gocms/app/validates/auth"
	"gocms/pkg/auth"
	"gocms/pkg/config"
	"gocms/pkg/logger"
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
	token := c.GetHeader("authorization")
	logger.Info("token", token)
	token = auth.ValidateToken(token)
	logger.Info("validate-token", token)

	if len(token) == 0 {
		response.ErrorResponse(401, "鉴权失败").WriteTo(c)
		return
	}

	user, err := jwtAction.ParseToken(token)
	if err != nil {
		response.ErrorResponse(401, err.Error()).WriteTo(c)
		return
	}

	response.SuccessResponse(user).WriteTo(c)
	return

}

func (*AuthController) Register(c *gin.Context) {

}
