package response

import (
	"github.com/gin-gonic/gin"
)

const (
	MaskNeedAuthor   = 8
	MaskParamMissing = 7
	StatusSuccess    = 0
)

type JsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data *interface{}) *JsonResponse {
	return &JsonResponse{
		Status:  StatusSuccess,
		Message: "success",
		Data:    data,
	}
}

func ErrorResponse(status int, message string) *JsonResponse {
	return &JsonResponse{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}

// 将 json 设为响应体.
// HTTP 状态码由应用状态码决定
func (that *JsonResponse) WriteTo(ctx *gin.Context) {
	code := 200
	if that.Status != StatusSuccess {
		code = that.responseCode()
	}
	ctx.JSON(code, that)
}

// 获取 HTTP 状态码. HTTP 状态码由 应用状态码映射
func (that *JsonResponse) responseCode() int {
	// todo 完善应用状态码对应 http 状态码
	if that.Status != StatusSuccess {
		return 500
	}
	return 200
}
