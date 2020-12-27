package admin

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/dgrijalva/jwt-go"
	"gocms/app/models/base"
	"gocms/app/models/permission"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
)

type Admin struct {
	base.BaseModel
	Account     string `json:"account"`
	Password    string `json:"password,omitempty"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Avatar      string `json:"avatar"`
}

// 此信息将写入鉴权中
type AuthAdmin struct {
	jwt.StandardClaims
	Account     string              `json:"account"`
	Description string              `json:"description,omitempty"`
	Email       string              `json:"email"`
	Phone       string              `json:"phone"`
	Roles       []string            `json:"roles"`
	Permissions []map[string]string `json:"permissions"`
}

var AuthUser AuthAdmin

func (Admin) TableName() string {
	return "admins"
}

// 获取安全认证的用户信息
func GetAuthAdmin(adminModel Admin) *AuthAdmin {

	r := &AuthAdmin{
		Account:     adminModel.Account,
		Description: adminModel.Description,
		Email:       adminModel.Email,
		Phone:       adminModel.Phone,
		Roles:       GetRolesForUser(adminModel.Account),
		Permissions: GetPermissionsForUser(adminModel.Account),
	}
	return r
}

// 生成权限
// 用于生成系统和全局的权限列表
// 更新创建动作
func GeneratePermissionNodes() {
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
			"method":     currentPermission[1],
			"permission": currentPermission[2],
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
