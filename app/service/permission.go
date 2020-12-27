package service

import (
	"gocms/app/models/base"
	"gocms/app/models/permission"
	"gocms/pkg/config"
	"gocms/pkg/help"
)

type PermissionService struct{}

func (*PermissionService) GetList(page int, pageSize int) *base.Result {
	admins := []permission.Permission{}
	offset := help.GetOffset(page, pageSize)
	total := 0

	config.Db.Model(&permission.Permission{}).Select("id, account, description, email, avatar, phone, created_at").Limit(pageSize).Offset(offset).Scan(&admins)
	config.Db.Model(&permission.Permission{}).Count(&total)

	data := base.Result{
		Page:     page,
		PageSize: pageSize,
		List:     admins,
		Total:    total,
	}

	return &data
}

// UpdateOrCreate 创建或者更新权限
func (*PermissionService) UpdateOrCreate(permissionModel permission.Permission) bool {
	if permissionModel.ID > 0 {
		return config.Db.Update(permissionModel).RowsAffected > 0
	}

	return config.Db.Create(permissionModel).RowsAffected > 0
}
