package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"gocms/app/http/admin/validates"
	"gocms/app/models/role"
	"gocms/app/service"
	"gocms/pkg/config"
	"gocms/pkg/response"
)

type RoleController struct{}

var rolenService = &service.RoleService{}

// 权限节点列表
func (that *RoleController) Index(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	list := rolenService.GetList(cast.ToInt(page), cast.ToInt(pageSize))

	response.SuccessResponse(list).WriteTo(c)
	return
}

// 数据保存
func (that *RoleController) Store(c *gin.Context) {
	var role role.RoleModel
	if !validates.VidateCreateOrUpdateRole(c, &role) {
		return
	}
	rolenService.UpdateOrCreateById(role)

	response.SuccessResponse().WriteTo(c)
	return
}

// 数据更新
func (that *RoleController) Save(c *gin.Context) {
	var params map[string]string
	params["id"] = c.Param("id")
	if !validates.VidateCreateOrUpdateRole(c, &params) {
		return
	}
	var model role.RoleModel
	_ = mapstructure.Decode(params, &model)
	rolenService.UpdateOrCreateById(model)

	response.SuccessResponse().WriteTo(c)
	return
}

// 数据删除
func (that *RoleController) destory(c *gin.Context) {
	id := c.Param("id")
	config.Db.Delete(role.RoleModel{}, "id = "+id)

	response.SuccessResponse().WriteTo(c)
	return
}
