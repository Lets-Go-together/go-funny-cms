package validate

import (
	"github.com/spf13/cast"
	"gocms/pkg/response"
	"gocms/wrap"
)

// 表示一个可验证对象, 实现该接口可以用于从 ctx 中取出接口参数, 并且在成功的情况自动写入到 params.
type Validatable interface {
	// 验证接口参数, 在错误情况应在该方法中处理错误
	//
	// @param	ctx 	包含需要验证参数请求上下文
	// @return	是否验证成功
	Validate(ctx *wrap.ContextWrapper) bool
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
func WithDefaultResponse(s interface{}, writer wrap.JSONWriter) bool {
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
func WithResponseMsg(s interface{}, writer wrap.JSONWriter, defaultMsg ...interface{}) bool {

	//msg := "参数验证失败，请检查"
	var msg string
	var success bool
	if len(defaultMsg) > 0 {
		msg = cast.ToString(defaultMsg[0])
		success, _ = Validate(s)
	} else {
		success, msg = Validate(s)
	}

	if success {
		return true
	}
	response.ErrorResponse(403, msg).WriteTo(writer)
	return false
}

// 验证 struct, 自动写入错误响应到 writer
//
// @param	s 		需要验证的结构体
// @param	writer	可输出 Json 响应对象
// @param 	status	应用状态码
// @return 验证是否成功
func WithResponse(s interface{}, status int, msg string, writer wrap.JSONWriter) bool {
	success, _ := Validate(s)
	if !success {
		response.ErrorResponse(status, msg).WriteTo(writer)
		return false
	}
	return true
}
