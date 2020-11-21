package response

import "github.com/gin-gonic/gin"

const (
	StatusSuccess = 0
)

type StatusCode struct {
	StatusSuccess int
	StatusError   int
	StatusRefused int
}

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

func (that *JsonResponse) WriteTo(ctx *gin.Context) {
	code := StatusSuccess
	if that.Status != StatusSuccess {
		code = that.responseCode()
	}
	ctx.JSON(code, that)
}

func (that *JsonResponse) responseCode() int {
	if that.Status != StatusSuccess {
		return 500
	}
	return 200
}
