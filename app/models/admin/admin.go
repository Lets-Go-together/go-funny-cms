package admin

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
	"gocms/app/models/base"
	"gocms/app/models/menu"
	"gocms/app/models/role"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/config"
)

type Admin struct {
	base.BaseModel
	Account     string            `validate:"required,gte=2,lte=8" json:"account"`
	Password    string            `json:"password,omitempty"`
	Description string            `validate:"required,gte=1,lte=60" json:"description"`
	Email       string            `validate:"required,email" json:"email"`
	Phone       string            `validate:"required" json:"phone"`
	Avatar      string            `validate:"required,url" json:"avatar"`
	Roles       []string          `json:"roles" gorm:"-"`
	RoleIds     base.IntJson      `json:"role_ids" gorm:"-"`
	Menus       []menu.MenuRouter `json:"menus" gorm:"-"`
}

// 此信息将写入鉴权中
type AuthAdmin struct {
	jwt.StandardClaims
	Account     string              `json:"account"`
	Description string              `json:"description,omitempty"`
	Email       string              `json:"email"`
	Phone       string              `json:"phone"`
	RoleIds     []int               `json:"role_ids" `
	Avatar      string              `json:"avatar"`
	Roles       []string            `json:"roles"`
	Menus       []menu.MenuRouter   `json:"menus",omitempty`
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
		Roles:       adminModel.Roles,
		Menus:       adminModel.Menus,
		Avatar:      adminModel.Avatar,
		Phone:       adminModel.Phone,
	}
	return r
}

// GetPermissions 通过用户账号获取权限
func GetPermissions(account string) []map[string]string {
	return rabc.GetPermissionsForUser(account)
}

// GetRoles 通过用户账号获取角色
func GetRoles(account string) []string {
	return rabc.GetRolesForUser(account)
}

// GetMenus 通过用户账号获取菜单
func GetMenus(roles []string, currentAccount string) []menu.MenuRouter {
	var RoleList []role.RoleModel
	var MenuList []menu.MenuRouter
	var menu_ids []int

	if currentAccount == config.Get("ACCOUNT") {
		config.Db.Model(menu.MenuModel{}).Order("weight desc").Scan(&MenuList)

		MenuList = GetMenuTreeRouter(MenuList, 1)
		return MenuList
	}

	config.Db.Model(role.RoleModel{}).Where("name in (?)", roles).Select("id, menu_ids").Scan(&RoleList)
	for _, role := range RoleList {
		menu_ids = append(menu_ids, role.MenuIds...)
	}

	config.Db.Model(menu.MenuModel{}).Where("id in (?)", menu_ids).Select("id, p_id").Scan(&MenuList)
	for _, item := range MenuList {
		if item.PId > 1 {
			menu_ids = append(menu_ids, item.PId)
		}
	}
	config.Db.Model(menu.MenuModel{}).Where("id in (?)", menu_ids).Order("weight desc").Scan(&MenuList)

	MenuList = GetMenuTreeRouter(MenuList, 1)
	return MenuList
}

func GetMenuTreeRouter(menus []menu.MenuRouter, pid int) []menu.MenuRouter {
	var list []menu.MenuRouter

	for _, v := range menus {
		if v.PId == pid {
			v.Children = GetMenuTreeRouter(menus, cast.ToInt(v.Id))
			list = append(list, v)
		}
	}

	return list
}
