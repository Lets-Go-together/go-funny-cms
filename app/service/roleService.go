package service

import (
	"gocms/app/models/base"
	"gocms/app/models/permission"
	"gocms/app/models/role"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
	"gocms/pkg/response"
	"gocms/wrap"
	"net/http"
	"sync"
)

type RoleService struct{}

var roleModel role.RoleModel

type RoleList struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	CreatedAt   string `json:"created_at"`
}

// 添加更新角色
func (*RoleService) UpdateOrCreateById(roleModel role.RoleModel, c *wrap.ContextWrapper) bool {
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
		config.Db.Model(&roleModel).Create(roleModel)
		_, err := config.Enforcer.AddRoleForUser("-", roleModel.Name)
		logger.PanicError(err, "添加角色", false)
	}

	currentPermission := roleModel.Permissions_ids
	if len(currentPermission) > 0 {
		updatePermissinForRole(currentPermission, roleModel.Name, originRoleModel.Name)
	}

	return true
}

func (*RoleService) GetList(page int, pageSize int, c *wrap.ContextWrapper) *base.Result {
	roles := []RoleList{}
	offset := help.GetOffset(page, pageSize)
	total := 0
	keyword := c.Query("keyword")

	query := config.Db.Model(&role.RoleModel{}).Select("id, name, description, status, created_at")
	if len(keyword) > 0 {
		query = query.Where("name like ?", "%"+keyword+"%")
	}

	query.Limit(pageSize).Offset(offset).Scan(&roles)
	query.Count(&total)

	data := base.Result{
		Page:     page,
		PageSize: pageSize,
		List:     roles,
		Total:    total,
	}

	return &data
}

// 更新角色权限
func updatePermissinForRole(permissionIds []int, roleAccount string, originRoleAccount string) {
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

// GetRolesName 通过id获取角色名称
func GetRolesName(ids []int) []string {
	roles := []role.RoleModel{}
	config.Db.Model(&roleModel).Select("id, name").Where("id in (?)", ids).Scan(&roles)

	names := []string{}
	for _, roleInfo := range roles {
		names = append(names, roleInfo.Name)
	}
	return names
}

// 根据所有权限和当前权限 获取权限节点ID
func (*RoleService) GetPermissionIdsByTree(currentPs []map[string]string, allPermissions []PermissionList) []int {
	var permission_ids []int
	var wg sync.WaitGroup

	for _, p := range allPermissions {
		for _, c := range currentPs {
			wg.Add(1)

			go func(p PermissionList, c map[string]string) {
				defer wg.Done()

				if p.Method == c["method"] && p.Url == c["permission"] {
					permission_ids = append(permission_ids, p.Id)
				}
			}(p, c)
		}
	}

	wg.Wait()

	return permission_ids
}
