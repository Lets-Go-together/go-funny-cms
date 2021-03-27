package controllers

import (
	"github.com/spf13/cast"
	"gocms/app/http/admin/validates"
	adminModel "gocms/app/models/admin"
	"gocms/app/service"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/config"
	"gocms/pkg/response"
	"gocms/wrap"
)

type AdminController struct{}

var adminService = &service.AdminService{}

// 管理员列表
func (*AdminController) Index(c *wrap.ContextWrapper) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	list := adminService.GetList(cast.ToInt(page), cast.ToInt(pageSize), c)

	response.SuccessResponse(list).WriteTo(c)
}

// 管理员创建
func (*AdminController) Store(c *wrap.ContextWrapper) {
	var params adminModel.Admin
	if !validates.VidateCreateOrUpdateAdmin(c, &params) {
		return
	}

	if adminService.UpdateOrCreate(params, c) {
		response.SuccessResponse().WriteTo(c)
	}
	return
}

// 管理员更新
func (*AdminController) Save(c *wrap.ContextWrapper) {
	var params adminModel.Admin
	params.ID = cast.ToUint64(c.Param("id"))
	if !validates.VidateCreateOrUpdateAdmin(c, &params) {
		return
	}

	adminService.UpdateOrCreate(params, c)
	response.SuccessResponse().WriteTo(c)
	return
}

// 数据删除
func (that *AdminController) Destory(c *wrap.ContextWrapper) {
	var param IdParam
	c.BindJSON(&param)

	config.Db.Delete(adminModel.Admin{}, "id = "+cast.ToString(param.Id))

	response.SuccessResponse().WriteTo(c)
	return
}

// 角色详情
func (that *AdminController) Show(c *wrap.ContextWrapper) {
	id := c.Param("id")
	result := adminModel.Admin{}
	config.Db.Model(adminModel.Admin{}).Omit("password").Select("id, account, description, email, phone, avatar, created_at, updated_at").Where("id = ?", id).First(&result)
	result.Roles = rabc.GetRolesForUser(result.Account)

	response.SuccessResponse(result).WriteTo(c)
	return
}
