package admin

import (
	"github.com/dgrijalva/jwt-go"
	"gocms/app/models/base"
	"gocms/pkg/auth/rabc"
)

type Admin struct {
	base.BaseModel
	Account     string   `validate:"required,gte=2,lte=8" json:"account"`
	Password    string   `json:"password,omitempty"`
	Description string   `validate:"required,gte=1,lte=60" json:"description"`
	Email       string   `validate:"required,email" json:"email"`
	Phone       string   `validate:"required" json:"phone"`
	Avatar      string   `validate:"required,url" json:"avatar"`
	Roles       []string `json:"roles" gorm:"-"`
	Role_ids    []int    `json:"role_ids,omitempty" gorm:"-"`
}

// 此信息将写入鉴权中
type AuthAdmin struct {
	jwt.StandardClaims
	Account     string              `json:"account"`
	Description string              `json:"description,omitempty"`
	Email       string              `json:"email"`
	Phone       string              `json:"phone"`
	Role_ids    []int               `json:"role_ids,omitempty" gorm:"-"`
	Avatar      string              `json:"avatar"`
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
		Roles:       adminModel.Roles,
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
