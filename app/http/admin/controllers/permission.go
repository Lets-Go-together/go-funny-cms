package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gocms/app/models/casbin"
	"gocms/app/models/permission"
	"gocms/app/service"
	"gocms/pkg/config"
	"gocms/pkg/response"
)

type PermissionController struct{}

var permissionService = &service.PermissionService{}

// 权限节点列表
func (that *PermissionController) Index(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	list := permissionService.GetList(cast.ToInt(page), cast.ToInt(pageSize))

	response.SuccessResponse(list).WriteTo(c)
	return
}

// Create 创建权限节点
func (that *PermissionController) Create(c *gin.Context) {

}

// 权限重置
func (that *PermissionController) reset(c *gin.Context) {
	// 是否权限清除并重置
	is_clear := cast.ToBool(c.PostForm("is_clear"))
	db := config.Db
	if is_clear == true {
		db.Delete(casbin.Casbin{}, "id > 0")
		db.Delete(permission.Permission{}, "id > 0")
	}

	// 自动创建权限
	permissionService.GeneratePermissionNodes()
	response.SuccessResponse().WriteTo(c)
	return
}
