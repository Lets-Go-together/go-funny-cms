package validate

import (
	"github.com/gin-gonic/gin"
	"gocms/pkg/response"
)

type Validatable interface {
	Validate(writer *gin.Context) string
}

// 验证器
// 返回验证器验证结果错误消息 和 bool (是否验证成功)
func Validate(s interface{}) (bool, string) {
	return CustomValidator.verify(s)
}

// 验证 structure,
func ValidateWithDefaultResponse(s interface{}, writer response.JSONWriter) bool {
	success, msg := Validate(s)
	if !success {
		response.ErrorResponse(403, msg).WriteTo(writer)
		return false
	}
	return true
}

func ValidateWithResponse(s interface{}, msg string, writer response.JSONWriter) bool {
	success, _ := Validate(s)
	if !success {
		response.ErrorResponse(403, msg).WriteTo(writer)
		return false
	}
	return true
}
