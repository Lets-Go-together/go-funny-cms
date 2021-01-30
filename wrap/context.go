package wrap

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONWriter interface {
	JSON(code int, data interface{})
	PureJSON(code int, data interface{})
	JSONP(code int, data interface{})
}

type Writer interface {
}

type Reader interface {
}

type HandlerFunc func(wrapper *ContextWrapper)

type ContextWrapper struct {
	*gin.Context

	// TODO 2021-1-25 屏蔽实际 Context, 便于控制修改需要暴露的接口和屏蔽不必要的接口, 利于规范代码
	//context *gin.Context
}

func (that *ContextWrapper) Accept(json interface{}) {
	that.JSON(http.StatusAccepted, &json)
}

func (that *ContextWrapper) Unauthorized() {
	that.String(http.StatusUnauthorized, "401 Unauthorized")
}

func (that *ContextWrapper) Forbidden() {
	that.String(http.StatusForbidden, "403 Forbidden")
}
