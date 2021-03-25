package controllers

import (
	"github.com/spf13/cast"
	"gocms/app/http/admin/validates"
	"gocms/app/models/casbin"
	"gocms/app/models/permission"
	"gocms/app/service"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/config"
	"gocms/pkg/response"
	"gocms/wrap"
)

type PermissionController struct{}

var permissionService = &service.PermissionService{}

// 权限节点列表
func (that *PermissionController) Index(c *wrap.ContextWrapper) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	list := permissionService.GetList(cast.ToInt(page), cast.ToInt(pageSize), c)

	response.SuccessResponse(list).WriteTo(c)
	return
}

// 数据保存
func (that *PermissionController) Store(c *wrap.ContextWrapper) {
	var permissionModel permission.Permission
	if !validates.VidateCreateOrUpdatePermission(c, &permissionModel) {
		return
	}

	//_ = mapstructure.Decode(params, &permissionModel)
	permissionService.UpdateOrCreate(permissionModel)

	response.SuccessResponse().WriteTo(c)
	return
}

// 数据更新
func (that *PermissionController) Save(c *wrap.ContextWrapper) {
	var permissionModel permission.Permission
	permissionModel.ID = cast.ToUint64(c.Param("id"))
	if !validates.VidateCreateOrUpdatePermission(c, &permissionModel) {
		return
	}

	permissionService.UpdateOrCreate(permissionModel)
	response.SuccessResponse().WriteTo(c)
	return
}

// 数据删除
func (that *PermissionController) Destory(c *wrap.ContextWrapper) {
	id := c.Param("id")
	config.Db.Delete(permission.Permission{}, "id = "+id)

	response.SuccessResponse().WriteTo(c)
	return
}

// 权限重置
func (that *PermissionController) reset(c *wrap.ContextWrapper) {
	// 是否权限清除并重置
	is_clear := cast.ToBool(c.PostForm("is_clear"))
	db := config.Db
	if is_clear == true {
		db.Delete(casbin.Casbin{}, "id > 0")
		db.Delete(permission.Permission{}, "id > 0")
	}

	// 自动创建权限
	rabc.GeneratePermissionNodes()
	response.SuccessResponse().WriteTo(c)
	return
}

// 角色详情
func (that *PermissionController) Show(c *wrap.ContextWrapper) {
	id := c.Param("id")
	type permissionInfo struct {
		Id     string
		Name   string `json:"name"`
		Icon   string `json:"icon"`
		Url    string `json:"url"`
		Status int    `json:"status"`
		Method string `json:"method"`
		Pid    int    `json:"pid"`
	}
	result := permissionInfo{}
	config.Db.Model(permission.Permission{}).Omit("").Where("id = ?", id).First(&result)

	response.SuccessResponse(result).WriteTo(c)
	return
}

// 获取权限节点树
func (that *PermissionController) Tree(c *wrap.ContextWrapper) {
	list := []service.PermissionList{
		{
			Id:       1,
			Name:     "根节点",
			Icon:     "link",
			Url:      "",
			Status:   1,
			Hidden:   2,
			Method:   "any",
			PId:      0,
			Children: nil,
		},
	}
	permissionsTree := permissionService.GetPermisstionTree()
	permissionsTree = append(list, permissionsTree...)

	response.SuccessResponse(permissionsTree).WriteTo(c)
}
