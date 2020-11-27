package validate

import (
	"github.com/gin-gonic/gin"
	"gocms/pkg/response"
)

// 表示一个可验证对象, 实现该接口可以用于从 ctx 中取出接口参数, 并且在成功的情况自动写入到 params.
type ValidationAction interface {
	// 验证接口参数, 在错误情况应在该方法中处理错误
	//
	// @param	ctx 	包含需要验证参数请求上下文
	// @param	params	用于该请求中的参数
	// @return	是否验证成功
	Validate(ctx *gin.Context, params interface{}) bool
}

// 验证器
// 返回验证器验证结果错误消息 和 bool (是否验证成功)
func Validate(s interface{}) (bool, string) {
	return CustomValidator.verify(s)
}

// 验证 struct, 自动写入默认错误响应到 writer
//
// @param	s 		需要验证的结构体
// @param	writer	可输出 Json 响应对象
// @return 验证是否成功
func WithDefaultResponse(s interface{}, writer response.JSONWriter) bool {
	success, msg := Validate(s)
	if !success {
		response.ErrorResponse(403, msg).WriteTo(writer)
		return false
	}
	return true
}

// 验证 struct, 自动写入错误响应到 writer
//
// @param	s 		需要验证的结构体
// @param 	msg		验证错误时自定义消息
// @param	writer	可输出 Json 响应对象
// @return 验证是否成功
func WithResponseMsg(s interface{}, msg string, writer response.JSONWriter) bool {
	success, _ := Validate(s)
	if !success {
		response.ErrorResponse(403, msg).WriteTo(writer)
		return false
	}
	return true
}

// 验证 struct, 自动写入错误响应到 writer
//
// @param	s 		需要验证的结构体
// @param	writer	可输出 Json 响应对象
// @param 	status	应用状态码
// @return 验证是否成功
func WithResponse(s interface{}, status int, msg string, writer response.JSONWriter) bool {
	success, _ := Validate(s)
	if !success {
		response.ErrorResponse(status, msg).WriteTo(writer)
		return false
	}
	return true
}
