package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gocms/app/http/admin/validates"
	"gocms/app/models"
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
	var params map[string]string
	if validates.VidateCreateOrUpdateAdmin(c, &params) == false {
		return
	}

	admin := models.Admin{
		Account:     params["account"],
		Password:    auth.CreatePassword(params["password"]),
		Description: params["description"],
		Email:       params["email"],
		Phone:       params["phone"],
		Avatar:      params["avatar"],
	}
	adminService.Create(admin)
	response.SuccessResponse().WriteTo(c)
	return
}

// 管理员更新
func (*AdminController) Update(c *gin.Context) {
	var params map[string]string
	id := cast.ToString(c.Param("id"))
	params["id"] = id
	if validates.VidateCreateOrUpdateAdmin(c, &params) == false {
		return
	}
	admin := models.Admin{
		Account:     params["account"],
		Password:    auth.CreatePassword(params["password"]),
		Description: params["description"],
		Email:       params["email"],
		Phone:       params["phone"],
		Avatar:      params["avatar"],
	}
	adminService.Update(admin, id)
	response.SuccessResponse().WriteTo(c)
	return
}
