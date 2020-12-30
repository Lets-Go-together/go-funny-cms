package admin

import (
	"github.com/dgrijalva/jwt-go"
	"gocms/app/models/base"
)

type Admin struct {
	base.BaseModel
	Account     string `validate:"required,gte=2,lte=8" json:"account"`
	Password    string `json:"password,omitempty"`
	Description string `validate:"required,gte=1,lte=60" json:"description"`
	Email       string `validate:"required,email" json:"email"`
	Phone       string `validate:"required" json:"phone"`
	Avatar      string `validate:"required,url" json:"avatar"`
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
	}
	return r
}
