package service

import (
	"github.com/gin-gonic/gin"
	"gocms/app/models/base"
	"gocms/app/models/permission"
	"gocms/app/models/role"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
	"gocms/pkg/response"
	"net/http"
)

type RoleService struct{}

var roleModel role.RoleModel

// 添加更新角色
func (*RoleService) UpdateOrCreateById(roleModel role.RoleModel, c *gin.Context) bool {
	var result bool
	var originRoleModel role.RoleModel

	if roleModel.ID > 0 {
		config.Db.Model(originRoleModel).Where("id = ?", roleModel.ID).First(&originRoleModel)
		if originRoleModel.ID == 0 {
			response.ErrorResponse(http.StatusNotFound, "角色未找到").WriteTo(c)
			return false
		}
		// 这里不用返回影响记录条数
		config.Db.Model(&roleModel).Where("id = ?", roleModel.ID).Update(roleModel)
	} else {
		result = config.Db.Model(&roleModel).Create(roleModel).RowsAffected > 0
		_, err := config.Enforcer.AddRoleForUser("-", roleModel.Name)
		logger.PanicError(err, "添加角色", false)
	}

	currentPermission := roleModel.Permissions
	if len(currentPermission) > 0 {
		updatePermissinForRole(currentPermission, roleModel.Name, originRoleModel.Name)
	}

	if !result {
		response.ErrorResponse(http.StatusInternalServerError, "更新异常").WriteTo(c)
		return false
	}
	return true
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
func updatePermissinForRole(permissionIds []interface{}, roleAccount string, originRoleAccount string) {
	permissionsModel := []permission.Permission{}
	config.Db.Model(permissionsModel).Where("id in (?)", permissionIds).Find(&permissionsModel)

	if len(originRoleAccount) > 0 {
		// 删除权限
		rabc.DeletePermissionsForUser(originRoleAccount)
	}

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
