package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gocms/app/models/admin"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"time"
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
func (*JwtAction) GetToken(userClaims *admin.AuthAdmin) string {
	expireAt := time.Duration(config.GetInt64("JWT_EXPIRE_AT", 60))
	userClaims.ExpiresAt = time.Now().Add(time.Minute * expireAt).Unix()
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

	if token.Valid {
		if _, ok := token.Claims.(*admin.AuthAdmin); ok && token.Valid {
			return userClaims, nil
		} else {
			return userClaims, err
		}
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("错误的token")
			logger.Info("错误的token", tokenString)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			logger.Info("token过期", tokenString)
		} else {
			logger.Info("无法处理这个token", tokenString)

		}
	}
	return userClaims, nil
}

// 刷新token， 参数同getToken'
func (*JwtAction) refreshToken() (token string, err error) {
	return "", nil
}
