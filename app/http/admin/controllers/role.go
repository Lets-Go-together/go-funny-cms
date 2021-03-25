package controllers

import (
	"github.com/spf13/cast"
	"gocms/app/http/admin/validates"
	"gocms/app/models/role"
	"gocms/app/service"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/config"
	"gocms/pkg/response"
	"gocms/wrap"
)

type RoleController struct{}

var rolenService = &service.RoleService{}

// 权限节点列表
func (that *RoleController) Index(c *wrap.ContextWrapper) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	list := rolenService.GetList(cast.ToInt(page), cast.ToInt(pageSize), c)

	response.SuccessResponse(list).WriteTo(c)
	return
}

// 数据保存
func (that *RoleController) Store(c *wrap.ContextWrapper) {
	var roleModel role.RoleModel
	if !validates.VidateCreateOrUpdateRole(c, &roleModel) {
		return
	}
	if rolenService.UpdateOrCreateById(roleModel, c) {
		response.SuccessResponse().WriteTo(c)
	}
	return
}

// 角色详情
func (that *RoleController) Show(c *wrap.ContextWrapper) {
	id := c.Param("id")
	result := role.RoleModel{}
	config.Db.Model(role.RoleModel{}).Where("id = ?", id).First(&result)
	result.Permissions = rabc.GetPermissionsForRole(result.Name)

	response.SuccessResponse(result).WriteTo(c)
	return
}

// 数据更新
func (that *RoleController) Save(c *wrap.ContextWrapper) {
	var roleModel role.RoleModel
	roleModel.ID = cast.ToUint64(c.Param("id"))
	if !validates.VidateCreateOrUpdateRole(c, &roleModel) {
		return
	}

	if rolenService.UpdateOrCreateById(roleModel, c) {
		response.SuccessResponse().WriteTo(c)
	}

	return
}

// 数据删除
func (that *RoleController) Destory(c *wrap.ContextWrapper) {
	id := c.Param("id")
	var roleModel role.RoleModel

	config.Db.Model(roleModel).Delete(role.RoleModel{}, "id = "+id)
	config.Db.Model(roleModel).Where("id = ?", id).First(&roleModel)
	rabc.DeletePermissionsForUser(roleModel.Name)
	rabc.DeleteRoleForUsers(roleModel.Name)

	response.SuccessResponse().WriteTo(c)
	return
}
