package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gocms/app/models/admin"
	"gocms/pkg/config"
)

var (
	signKey []byte
)

type JwtAction struct{}

func init() {
	key := config.GetString("JWT_SIGN")
	signKey = []byte(key)
}

// 获取token
// 必须传参 需要保存的用户信息
func (*JwtAction) GetToken(userClaims admin.AuthAdmin) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, _ := token.SignedString(signKey)
	return tokenString
}

// token 解析
// 返回用户模型的核型信息
// 注意: 第二个参数不为 nil 的时候，则表示解析失败
func (*JwtAction) ParseToken(tokenString string) (admin.AuthAdmin, error) {
	userClaims := admin.AuthAdmin{}
	token, err := jwt.ParseWithClaims(tokenString, &userClaims, func(token *jwt.Token) (i interface{}, e error) {
		return signKey, e
	})

	if err != nil {
		return userClaims, err
	}

	if _, ok := token.Claims.(*admin.AuthAdmin); ok && token.Valid {
		fmt.Println(userClaims)
		return userClaims, nil
	} else {
		return userClaims, err
	}
}

// 刷新token， 参数同getToken'
func (*JwtAction) refreshToken() (token string, err error) {
	return "", nil
}
