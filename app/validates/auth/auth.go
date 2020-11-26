package auth

import (
	"github.com/gin-gonic/gin"
	"gocms/app/validates/validate"
)

var LoginData *LoginParams

// 登陆需要的参数
type LoginParams struct {
	Account  string `validate:"required" json:"account"`
	Password string `validate:"required" json:"password"`
}

type LoginAction struct{}

func (*LoginAction) Validate(c *gin.Context) string {
	LoginData = &LoginParams{
		Account:  c.PostForm("account"),
		Password: c.PostForm("password"),
	}

	if isSuccess, msg := validate.Validate(LoginData); !isSuccess {
		return msg
	}
	return ""
}

func (*LoginAction) GetLoginData() *LoginParams {
	return LoginData
}
