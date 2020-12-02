package controllers

import (
	"github.com/gin-gonic/gin"
	"gocms/app/models/admin"
	"gocms/app/models/model"
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

	var params authValidate.LoginParams
	if !authValidateAction.Validate(c, &params) {
		return
	}

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

// 注册
func (*AuthController) Register(c *gin.Context) {
	action := authValidate.RegisterAction{}
	var params authValidate.RegisterParams
	if !action.Validate(c, &params) {
		return
	}

	e := admin.Admin{}
	exist := config.Db.Where("account = ?", params.Account).First(&e).RowsAffected > 0
	if exist {
		response.ErrorResponse(1002, "用户名已存在").WriteTo(c)
		return
	}

	exist = config.Db.Where("email = ?", params.Account).First(&e).RowsAffected > 0
	if exist {
		response.ErrorResponse(1002, "邮箱已注册").WriteTo(c)
		return
	}

	password := auth.CreatePassword(params.Password) // salt(params.password)

	newAdmin := admin.Admin{
		BaseModel:   model.BaseModel{},
		Account:     params.Account,
		Password:    password,
		Description: "",
		Email:       params.Email,
		Phone:       "",
		Avatar:      "",
	}
	config.Db.NewRecord(&newAdmin)
	response.SuccessResponse("").WriteTo(c)
}
