package auth

import (
	"github.com/gin-gonic/gin"
	"gocms/app/validates/validate"
	"gocms/pkg/logger"
)

// 登陆需要的参数
type LoginParams struct {
	Account  string `validate:"required" json:"account"`
	Password string `validate:"required" json:"password"`
}

type LoginAction struct {
}

func (*LoginAction) Validate(c *gin.Context, params interface{}) bool {
	params = LoginParams{
		Account:  c.PostForm("account"),
		Password: c.PostForm("password"),
	}

	// 通过JSON获取数据
	//err := c.BindJSON(params)
	//if err != nil {
	//	fmt.Println(params)
	//	logger.PanicError(err, "登录参数验证", false)
	//	response.ErrorResponse(http.StatusForbidden, err.Error()).WriteTo(c)
	//	return false
	//}

	// 自动写入默认校验错误消息和状态码到错误响应
	//return validate.WithDefaultResponse(params, c)

	// 自动写入自定义消息到错误响应
	//return validate.WithResponse(c, "请检查账号与密码是否正确", c)

	// 自定义错误码和消息
	return validate.WithResponse(params, 403, "Error", c)
}

// --- 注册相关

type RegisterParams struct {
	Account  string `validate:"username" json:"account"`
	Password string `validate:"required" json:"password"`
	Email    string `validate:"email" json:"email"`
	Captcha  string `validate:"alphanumeric,len=5" json:"captcha"`
}

type RegisterAction struct {
}

func (that *RegisterAction) Validate(c *gin.Context, params interface{}) bool {
	err := c.BindJSON(params)
	if err != nil {
		logger.PanicError(err, "注册参数验证", false)
		return false
	}
	return validate.WithResponse(params, 403, "请检查参数", c)
}

// --- 用户创建、更新相关
type EditParams struct {
	Account  string `validate:"username" json:"account"`
	Password string `validate:"required" json:"password"`
	Phone    string `validate:"required" json:"Phone"`
	Email    string `validate:"email" json:"email"`
}
