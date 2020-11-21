package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var (
	signKey []byte
)

type JwtAction struct{}

type UserInfo struct {
	Name string
}

type UserClaims struct {
	*jwt.StandardClaims
	UserInfo
}

func init() {
	key := "123455sign"
	signKey = []byte(key)
}

func (*JwtAction) GetToken() string {
	token, err := createToken("chenf")
	if err != nil {

	}

	return token
}

func (*JwtAction) ParseToken(tokenString string) string {
	token, e := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return signKey, nil
	})
	if e != nil {
		return "error-"
	}

	fmt.Println(token)

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Name, claims.StandardClaims.ExpiresAt)
		return claims.Name
	} else {
		return e.Error()
	}
}

// 创建Token
func createToken(user string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))

	// 配置
	t.Claims = &UserClaims{
		&jwt.StandardClaims{
			ExpiresAt: 0,
		},
		UserInfo{Name: user},
	}

	return t.SignedString(signKey)
}
