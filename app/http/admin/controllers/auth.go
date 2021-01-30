package controllers

import (
	"gocms/app/http/admin/validates"
	"gocms/app/models/admin"
	"gocms/app/models/base"
	"gocms/pkg/auth"
	"gocms/pkg/config"
	"gocms/pkg/response"
	"gocms/wrap"
	"net/http"
)

type AuthController struct{}

var (
	jwtAction = auth.JwtAction{}
)

func (*AuthController) Login(c *wrap.ContextWrapper) {
	authValidateAction := validates.LoginAction{}

	params := validates.LoginParams{}
	if !authValidateAction.Validate(c, &params) {
		return
	}

	adminModel := admin.Admin{}
	config.Db.Where("account = ?", params.Account).Find(&adminModel)

	if adminModel.ID == 0 {
		response.ErrorResponse(http.StatusForbidden, "用户不存在").WriteTo(c)
		return
	}

	password := adminModel.Password
	if !auth.ValidatePassword(password, params.Password) {
		response.ErrorResponse(http.StatusForbidden, "密码错误").WriteTo(c)
		return
	}

	authAdmin := admin.GetAuthAdmin(adminModel)
	token := jwtAction.GetToken(authAdmin)

	response.SuccessResponse(map[string]string{
		"token": token,
	}).WriteTo(c)

	return
}

// 我的信息
func (*AuthController) Me(c *wrap.ContextWrapper) {
	user := admin.AuthUser
	response.SuccessResponse(user).WriteTo(c)
	return
}

// 注销
func (*AuthController) Logout(c *wrap.ContextWrapper) {
	user := admin.AuthUser
	response.SuccessResponse(user).WriteTo(c)
}

func (*AuthController) Register2(c *wrap.ContextWrapper, params *validates.RegisterParams) {
	// CRUD
	// ...
	response.SuccessResponse("user register successful: " + params.Account).WriteTo(c)
}

// 注册
func (*AuthController) Register(c *wrap.ContextWrapper) {
	action := validates.RegisterAction{}
	var params validates.RegisterParams
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
		BaseModel:   base.BaseModel{},
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
