package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gocms/app/models/users"
)

var (
	signKey []byte
)

type JwtAction struct{}

type UserClaims struct {
	jwt.StandardClaims
	authUser struct {
		users.AuthUser
	}
}

func init() {
	key := "123455sign"
	signKey = []byte(key)
}

func (*JwtAction) GetToken(user users.AuthUser) string {
	claims := &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 0,
		},
		authUser: struct {
			user
		}{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(signKey)
	return tokenString
}

func (*JwtAction) ParseToken(tokenString string) (string, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return signKey, e
	})

	if err != nil {
		return err.Error()
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		//return claims.Name
	} else {
		return err.Error()
	}
}
