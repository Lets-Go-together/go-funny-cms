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

type roleController struct{}

var rolenService = &service.RoleService{}

// 权限节点列表
func (that *roleController) Index(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	list := rolenService.GetList(cast.ToInt(page), cast.ToInt(pageSize))

	response.SuccessResponse(list).WriteTo(c)
	return
}

// 数据保存
func (that *roleController) Store(c *gin.Context) {
	var params map[string]string
	if validates.VidateCreateOrUpdateRole(c, &params) == false {
		return
	}
	var model role.RoleModel
	_ = mapstructure.Decode(params, &model)

	response.SuccessResponse().WriteTo(c)
	return
}

// 数据更新
func (that *roleController) Save(c *gin.Context) {
	var params map[string]string
	params["id"] = c.Param("id")
	if validates.VidateCreateOrUpdateRole(c, &params) == false {
		return
	}
	var model role.RoleModel
	_ = mapstructure.Decode(params, &model)

	response.SuccessResponse().WriteTo(c)
	return
}

// 数据删除
func (that *roleController) destory(c *gin.Context) {
	id := c.Param("id")
	config.Db.Delete(role.RoleModel{}, "id = "+id)

	response.SuccessResponse().WriteTo(c)
	return
}
