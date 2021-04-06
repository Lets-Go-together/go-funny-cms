package rabc

import (
	"github.com/cheggaaa/pb/v3"
	"gocms/app/models/permission"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
)

// 生成权限
// 用于生成系统和全局的权限列表
// 更新创建动作
func GeneratePermissionNodes() {
	routers := config.GetAllRoutes()
	db := config.Db
	count := len(routers)
	bar := pb.StartNew(count)

	//rootNode := &permission.Permission{
	//	Name:   "根节点",
	//	Icon:   "link",
	//	Url:    "",
	//	Status: 1,
	//	Method: "any",
	//	PId:    0,
	//}
	//db.Model(&permission.Permission{}).Omit("pid", "status", "icon", "hidden").Create(&rootNode)

	for _, router := range routers {
		permissionModel := &permission.Permission{}
		routerCopy := &permission.Permission{
			Url:    router["url"],
			Method: router["method"],
		}

		// 忽略需要过滤的节点
		if !Filter(routerCopy) {
			bar.Increment()
			continue
		}

		db.Where(routerCopy).Select("id").First(permissionModel)

		//routerCopy.PId = cast.ToInt(rootNode.ID)
		if permissionModel.ID == 0 {
			routerCopy.Name = router["name"]
			db.Model(&permission.Permission{}).Omit("pid", "status", "icon", "hidden").Create(routerCopy)
		} else {
			db.Model(&permission.Permission{}).Where("id = ?", permissionModel.ID).Update(routerCopy)
		}
		bar.Increment()
	}
}

// GetPermissionNodes 获取全部权限节点
// 参数为空的的时候：返回本地的路由中的权限节点
// 否则返回数据库的
func GetPermissionNodes(params ...interface{}) []map[string]string {
	if param := help.GetDefaultParam(params); param == nil {
		return config.GetAllRoutes()
	}
	return []map[string]string{}
}

// HasPermissionForUser 检查是否有这个权限
// 参数必须存在
func HasPermissionForUser(account string, permission string, method string) bool {
	e := config.Enforcer
	roles, _ := e.GetRolesForUser(account)
	for _, role := range roles {
		ok := config.Enforcer.HasPermissionForUser(role, permission, method)
		if ok {
			return ok
		}
	}
	return false
}

// AddRoleForUser 授权用户到角色
func AddRoleForUser(role string, account string) bool {
	ok, err := config.Enforcer.AddRoleForUser(account, role)
	logger.PanicError(err, "授权用户到角色", false)
	return ok
}

// AddRoleForUser 批量授权用户到角色
func AddRolesForUser(account string, role []string) bool {
	ok, err := config.Enforcer.AddRolesForUser(account, role)
	logger.PanicError(err, "授权用户到角色", false)
	return ok
}

// AddPermissionForUser 添加权限到角色
func AddPermissionForUser(permission string, method string, role string) bool {
	ok, err := config.Enforcer.AddPermissionForUser(role, permission, method)
	logger.PanicError(err, "添加权限到角色", false)
	return ok
}

// GetPermissionsForRole 获取用户（角色）权限
func GetPermissionsForRole(role string) []map[string]string {
	var permissions []map[string]string
	currentPermissions := config.Enforcer.GetPermissionsForUser(role)
	for _, currentPermission := range currentPermissions {
		permissions = append(permissions, map[string]string{
			"method":     currentPermission[2],
			"permission": currentPermission[1],
		})
	}
	return permissions
}

// GetRolesForUser 获取用户角色
func GetRolesForUser(user string) []string {
	roles, err := config.Enforcer.GetRolesForUser(user)
	logger.PanicError(err, "获取用户角色", false)
	return roles
}

// GetPermissionsForUser 获取用户权限
func GetPermissionsForUser(account string) []map[string]string {
	var permissions []map[string]string
	roles := GetRolesForUser(account)
	for _, role := range roles {
		rolePermissions := GetPermissionsForRole(role)
		for _, rolePermission := range rolePermissions {
			permissions = append(permissions, rolePermission)
		}
	}
	return permissions
}

// DeletePermissionsForUser 删除角色的权限
func DeletePermissionsForUser(role string) bool {
	_, err := config.Enforcer.DeletePermissionsForUser(role)
	logger.PanicError(err, "删除角色权限", false)
	return true
}

// DeletePermissionsForUser 删除拥有对应角色的(用户角色权限)
func DeleteRoleForUsers(role string) bool {
	users, err := config.Enforcer.GetUsersForRole(role)
	logger.PanicError(err, "获取具有角色的用户", false)
	for _, user := range users {
		_, _ = config.Enforcer.DeleteRoleForUser(user, role)
	}
	return true
}

// DeleteRolesForUser 删除用户的所有角色
func DeleteRolesForUser(account string) bool {
	_, err := config.Enforcer.DeleteRolesForUser(account)
	logger.PanicError(err, "删除用户的所有角色", false)
	return true
}
