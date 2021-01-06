package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gocms/app/http/admin/validates"
	adminModel "gocms/app/models/admin"
	"gocms/app/service"
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
func (*AdminController) Store(c *gin.Context) {
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
func (*AdminController) Save(c *gin.Context) {
	var params adminModel.Admin
	params.ID = cast.ToUint64(c.Param("id"))
	if !validates.VidateCreateOrUpdateAdmin(c, &params) {
		return
	}

	adminService.UpdateOrCreate(params, c)
	response.SuccessResponse().WriteTo(c)
	return
}
