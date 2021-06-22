package wrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

/**
把 Context 的方法按用途抽象成若干接口, 在具体使用其中部分方法的时候应按接口引用代替使用 Context 引用.
用接口分离抽象了具体的功能, 隔离了不必要的功能. 例如在只需要读取的地方只传入 Reader 作为参数更加明显的说明了这个方法只需要进行读取操作.
*/

type JSONWriter interface {
	JSON(code int, data interface{})
}

type QueryReader interface {
	Query(key string) string
	DefaultQuery(key string, def string) string
}

type PathParamReader interface {
	Param(key string) string
}

type Writer interface {
	JSONWriter
}

type Reader interface {
	BindJSON(i interface{}) error
	ShouldBind(i interface{}) error
	PostForm(key string) string
}

type HandlerFunc func(wrapper *ContextWrapper)

// 包装一下 gin.Context, 方便控制接口的暴露和实现, 以及后期拓展一些功能.
type ContextWrapper struct {
	//*gin.Context

	// TODO 2021-1-25 屏蔽实际 Context, 便于控制修改需要暴露的接口和屏蔽不必要的接口, 利于规范代码
	Ctx *gin.Context
}

func Context(Ctx *gin.Context) *ContextWrapper {
	return &ContextWrapper{
		Ctx: Ctx,
	}
}

func (that *ContextWrapper) ResponseJson(json interface{}) {
	that.Ctx.JSON(http.StatusAccepted, &json)
}

func (that *ContextWrapper) ResponseString(str string) {
	that.Ctx.String(http.StatusAccepted, str)
}

func (that *ContextWrapper) Unauthorized() {
	that.Ctx.AbortWithStatus(http.StatusUnauthorized)
}

func (that *ContextWrapper) Forbidden() {
	that.Ctx.AbortWithStatus(http.StatusForbidden)
}

func (that *ContextWrapper) Query(key string) string {
	return that.Ctx.Query(key)
}

func (that *ContextWrapper) DefaultQuery(key string, def interface{}) string {
	ret := that.Ctx.DefaultQuery(key, cast.ToString(def))
	if ret == "" {
		return fmt.Sprintf("%s", def)
	}
	return ret
}

func (that *ContextWrapper) Param(key string) string {
	return that.Ctx.Param(key)
}

func (that *ContextWrapper) JSON(code int, data interface{}) {
	that.Ctx.JSON(code, data)
}

func (that *ContextWrapper) BindJSON(i interface{}) error {
	return that.Ctx.BindJSON(i)
}

func (that *ContextWrapper) ShouldBind(i interface{}) error {
	return that.Ctx.ShouldBind(i)
}

func (that *ContextWrapper) ShouldBindQuery(i interface{}) error {
	return that.Ctx.ShouldBindQuery(i)
}

func (that *ContextWrapper) GetQueryArray(key string) ([]string, bool) {
	return that.Ctx.GetQueryArray(key)
}

func (that *ContextWrapper) PostForm(key string) string {
	return that.Ctx.PostForm(key)
}
