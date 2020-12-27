package service

import (
	"gocms/app/models/base"
	"gocms/app/models/role"
	"gocms/pkg/config"
	"gocms/pkg/help"
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
