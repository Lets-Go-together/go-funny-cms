package admin

import (
	"github.com/dgrijalva/jwt-go"
	"gocms/app/models/base"
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
	Account     string   `json:"account"`
	Description string   `json:"description,omitempty"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Roles       []string `json:"roles"`
}

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
		Roles: []string{
			"admin",
		},
	}
	return r
}
