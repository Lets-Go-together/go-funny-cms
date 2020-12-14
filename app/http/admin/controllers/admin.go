package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gocms/app/http/admin/validates"
	adminModel "gocms/app/models/admin"
	"gocms/app/service"
	"gocms/pkg/auth"
	"gocms/pkg/response"
)

type AdminController struct{}

var adminService = &service.AdminService{}

// 管理员列表
func (*AdminController) Index(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	list := adminService.GetList(cast.ToInt(page), cast.ToInt(pageSize))

	response.SuccessResponse(list).WriteTo(c)
}

// 管理员创建
func (*AdminController) Create(c *gin.Context) {
	var params validates.Admin
	if validates.VidateCreateAdmin(c, &params) == false {
		return
	}
	admin := adminModel.Admin{
		Account:     params.Account,
		Password:    auth.CreatePassword(params.Password),
		Description: params.Description,
		Email:       params.Email,
		Phone:       params.Phone,
		Avatar:      params.Avatar,
	}
	adminService.Create(admin)
	response.SuccessResponse().WriteTo(c)
	return
}

// 管理员更新
func (*AdminController) Update(c *gin.Context) {
	params := validates.AdminUpdate{
		ID: cast.ToUint64(c.Param("id")),
	}
	if validates.VidateUpdateAdmin(c, &params) == false {
		return
	}
	admin := adminModel.Admin{
		Account:     params.Account,
		Password:    auth.CreatePassword(params.Password),
		Description: params.Description,
		Email:       params.Email,
		Phone:       params.Phone,
		Avatar:      params.Avatar,
	}
	adminService.Update(admin, params.ID)
	response.SuccessResponse().WriteTo(c)
	return
}
