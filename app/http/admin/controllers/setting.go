package controllers

import (
	"fmt"
	"gocms/app/http/admin/validates"
	"gocms/app/models/admin"
	"gocms/app/service"
	"gocms/app/validates/validate"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/response"
	"gocms/wrap"
)

type SettingController struct{}

// bindEmail 账号绑定Email
func (s SettingController) bindEmail(c *wrap.ContextWrapper) {
	var params validates.EmailValidate
	c.ShouldBind(&params)

	if !validate.WithResponseMsg(params, c, "Email 参数错误") {
		return
	}

	adminUser := admin.Admin{}
	config.Db.Model(&admin.Admin{}).Where("email", params.Email).Select("email").Find(&adminUser)
	if len(adminUser.Email) > 0 {
		response.ErrorResponse(401, "Email 已被占用")
		return
	}

	token, err := help.Enctrypt(fmt.Sprintf("%s,%s", params.Email, admin.AuthUser.Id))
	if err != nil {
		response.ErrorResponse(500, "服务错误")
		return
	}

	setting := new(service.SettingService)
	setting.ToBindEmail(token)

	response.SuccessResponse("您将收到一封确认邮件，请查收！")
	return
}
