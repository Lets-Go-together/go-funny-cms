package service

import (
	"gocms/app/models/base"
	"gocms/app/models/permission"
	"gocms/pkg/config"
	"gocms/pkg/help"
)

type PermissionService struct{}
type PermissionList struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Icon      string      `json:"icon"`
	Url       string      `json:"url"`
	Status    int         `json:"status"`
	Method    string      `json:"method"`
	Pid       int         `json:"pid"`
	CreatedAt base.TimeAt `json:"created_at"`
}

func (*PermissionService) GetList(page int, pageSize int) *base.Result {
	var permissions []PermissionList
	offset := help.GetOffset(page, pageSize)
	total := 0

	config.Db.Model(&permission.Permission{}).Select("id, name, icon, url, status, method, p_id, hidden, created_at").Limit(pageSize).Offset(offset).Scan(&permissions)
	config.Db.Model(&permission.Permission{}).Count(&total)

	data := base.Result{
		Page:     page,
		PageSize: pageSize,
		List:     permissions,
		Total:    total,
	}

	return &data
}

// UpdateOrCreate 创建或者更新权限
func (*PermissionService) UpdateOrCreate(permissionModel permission.Permission) bool {
	if permissionModel.ID > 0 {
		return config.Db.Model(permissionModel).Update(permissionModel).RowsAffected > 0
	}

	return config.Db.Model(permissionModel).Create(permissionModel).RowsAffected > 0
}
