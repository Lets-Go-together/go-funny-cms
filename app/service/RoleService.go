package service

import (
	"gocms/app/models/role"
	"gocms/pkg/config"
	"gocms/pkg/logger"
)

type RoleService struct{}

var roleModel role.RoleModel

// 添加更新角色
func (*RoleService) UpdateOrCreateById(role role.RoleModel) bool {
	if role.ID > 0 {
		return config.Db.Model(&roleModel).Where("id = ?", role.ID).Update(role).RowsAffected > 0
	}

	config.Db.Model(&roleModel).Create(role)
	ok, err := config.Enforcer.AddRoleForUser("-", role.Name)
	logger.PanicError(err, "添加角色", false)
	return ok
}
