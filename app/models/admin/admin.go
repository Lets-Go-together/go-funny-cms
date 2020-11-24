package admin

import (
	"github.com/dgrijalva/jwt-go"
	"gocms/app/models/model"
)

type Admin struct {
	model.BaseModel
	Account     string `json:"account"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

// 此信息将写入鉴权中
type AuthAdmin struct {
	jwt.StandardClaims
	Admin
}
