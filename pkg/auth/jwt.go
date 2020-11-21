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

func init() {
	key := "AllYourBase"
	signKey = []byte(key)
}

func (*JwtAction) GetToken(userClaims users.AuthUser) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, _ := token.SignedString(signKey)
	return tokenString
}

func (*JwtAction) ParseToken(tokenString string) (users.AuthUser, error) {

	userClaims := users.AuthUser{}
	token, err := jwt.ParseWithClaims(tokenString, &userClaims, func(token *jwt.Token) (i interface{}, e error) {
		return signKey, e
	})

	if err != nil {
		return userClaims, err
	}

	if _, ok := token.Claims.(*users.AuthUser); ok && token.Valid {
		fmt.Println(userClaims)
		return userClaims, nil
	} else {
		return userClaims, err
	}
}
