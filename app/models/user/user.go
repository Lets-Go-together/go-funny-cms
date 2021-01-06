package user

import (
	"github.com/dgrijalva/jwt-go"
)

type UserModel struct{}

// 此信息将写入鉴权中
type AuthUser struct {
	jwt.StandardClaims
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (AuthUser) TableName() string {
	return "users"
}
