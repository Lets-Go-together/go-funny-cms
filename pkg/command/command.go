package command

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli/v2"
	"gocms/app/models/users"
	"gocms/bootstrap"
	"gocms/pkg/auth"
)

func InitApp() *cli.App {
	app := &cli.App{
		Name:  "Start server",
		Usage: "--",
		Action: func(c *cli.Context) error {
			AppServer()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "process",
				Aliases: []string{"s"},
				Usage:   "可以在这里触发一些自定义脚本",
				Action: func(c *cli.Context) error {
					jwtTest()
					return nil
				},
			},
		},
	}

	return app
}

// 服务器
func AppServer() {
	bootstrap.SteupRoute()
}

func jwtTest() {

	action := new(auth.JwtAction)
	tokenString := action.GetToken(users.AuthUser{
		Id:   10,
		Name: "chenf",
	})

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	token, e := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(e, "e")
	}

}
