package service

import (
	"gocms/app/models/base"
	"gocms/app/models/permission"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"

	"github.com/cheggaaa/pb/v3"
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

// 生成权限
// 用于生成系统和全局的权限列表
// 更新创建动作
func (*PermissionService) GeneratePermissionNodes() {
	routers := config.GetAllRoutes()
	db := config.Db
	count := len(routers)
	bar := pb.StartNew(count)
	for _, router := range routers {
		permissionModel := &permission.Permission{}
		routerCopy := &permission.Permission{
			Url:    router["url"],
			Method: router["method"],
		}
		db.Where(routerCopy).Select("id").First(permissionModel)
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
func (*PermissionService) GetPermissionNodes(params ...interface{}) []map[string]string {
	if param := help.GetDefaultParam(params); param == nil {
		return config.GetAllRoutes()
	}
	return []map[string]string{}
}

// HasPermissionForUser 检查是否有这个权限
// 参数必须存在
func (*PermissionService) HasPermissionForUser(account string, permission string, method string) bool {
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
func (*PermissionService) AddRoleForUser(role string, account string) bool {
	ok, err := config.Enforcer.AddRoleForUser(account, role)
	logger.PanicError(err, "授权用户到角色", false)
	return ok
}

// AddPermissionForUser 添加权限到角色
func (*PermissionService) AddPermissionForUser(permission string, method string, role string) bool {
	ok, err := config.Enforcer.AddPermissionForUser(role, permission, method)
	logger.PanicError(err, "添加权限到角色", false)
	return ok
}

// GetPermissionsForRole 获取用户（角色）权限
func (*PermissionService) GetPermissionsForRole(role string) []map[string]string {
	var permissions []map[string]string
	currentPermissions := config.Enforcer.GetPermissionsForUser(role)
	for _, currentPermission := range currentPermissions {
		permissions = append(permissions, map[string]string{
			"method":     currentPermission[1],
			"permission": currentPermission[2],
		})
	}
	return permissions
}

// GetRolesForUser 获取用户角色
func (that *PermissionService) GetRolesForUser(user string) []string {
	roles, err := config.Enforcer.GetRolesForUser(user)
	logger.PanicError(err, "获取用户角色", false)
	return roles
}

// GetPermissionsForUser 获取用户权限
func (that *PermissionService) GetPermissionsForUser(account string) []map[string]string {
	var permissions []map[string]string
	roles := that.GetRolesForUser(account)
	for _, role := range roles {
		rolePermissions := that.GetPermissionsForRole(role)
		for _, rolePermission := range rolePermissions {
			permissions = append(permissions, rolePermission)
		}
	}
	return permissions
}
