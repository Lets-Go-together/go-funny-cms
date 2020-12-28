package service

import (
	"gocms/app/models/base"
	"gocms/app/models/permission"
	"gocms/app/models/role"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
)

type RoleService struct{}

var roleModel role.RoleModel

// 添加更新角色
func (*RoleService) UpdateOrCreateById(role role.RoleModel) bool {
	var result bool
	if role.ID > 0 {
		result = config.Db.Model(&roleModel).Where("id = ?", role.ID).Update(role).RowsAffected > 0
	} else {
		result = config.Db.Model(&roleModel).Create(role).RowsAffected > 0
		_, err := config.Enforcer.AddRoleForUser("-", role.Name)
		logger.PanicError(err, "添加角色", false)
	}

	currentPermission := roleModel.Permissions
	if len(currentPermission) > 0 {
		updatePermissinForRole(currentPermission, roleModel.Name)
	}

	return result
}

func (*RoleService) GetList(page int, pageSize int) *base.Result {
	admins := []role.RoleModel{}
	offset := help.GetOffset(page, pageSize)
	total := 0

	config.Db.Model(&role.RoleModel{}).Select("id, account, description, email, avatar, phone, created_at").Limit(pageSize).Offset(offset).Scan(&admins)
	config.Db.Model(&role.RoleModel{}).Count(&total)

	data := base.Result{
		Page:     page,
		PageSize: pageSize,
		List:     admins,
		Total:    total,
	}

	return &data
}

// 更新角色权限
func updatePermissinForRole(permissionIds []string, roleAccount string) {
	permissionsModel := []permission.Permission{}
	config.Db.Where("id in (?)", permissionIds).Find(&permissionsModel)

	// 删除权限
	rabc.DeletePermissionsForUser(roleAccount)

	// 赋值权限
	for _, permissionModel := range permissionsModel {
		rabc.AddPermissionForUser(permissionModel.Url, permissionModel.Method, roleAccount)
	}

	logger.PanicError(nil, "权限更新: "+roleAccount, false)
}

// UpdateOrCreate 创建或者更新权限
func (*RoleService) UpdateOrCreate(roleModel role.RoleModel) bool {
	if roleModel.ID > 0 {
		return config.Db.Update(roleModel).RowsAffected > 0
	}

	return config.Db.Create(roleModel).RowsAffected > 0
}
