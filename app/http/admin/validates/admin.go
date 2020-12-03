package validates

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gocms/app/validates/validate"
	"gocms/pkg/logger"
)

// 验证管理员创建参数
func VidateCreateAdmin(c *gin.Context) bool {
	type Admin struct {
		Description string `validate:"required" json:"description"`
		Email       string `validate:"required" json:"email"`
		Phone       string `validate:"required" json:"phone"`
		Avatar      string `validate:"required" json:"avatar"`
	}
	params := &Admin{}
	err := c.ShouldBindJSON(&params)

	fmt.Println(params)
	if err != nil {
		logger.PanicError(err, "登录参数验证", false)
	}

	return validate.WithResponseMsg(params, c)
}
