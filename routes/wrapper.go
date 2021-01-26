package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gocms/app/validates/validate"
	"net/http"
	"reflect"
)

type Controller interface {
}

type acceptor func(ctx *gin.Context, params interface{})

var typeGinContext = reflect.TypeOf((*gin.Context)(nil))
var typeValidatable = reflect.TypeOf((*validate.ValidationAction)(nil)).Elem()

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

type iRoute interface {
	setup(parent gin.IRouter)
}

type route struct {
	method       string
	relativePath string
	handlerFunc  interface{}
}

type routeGroup struct {
	name   string
	routes []iRoute
}

type routeUse struct {
	middleware []gin.HandlerFunc
	routes     []iRoute
}

func setupRoutes(rt iRoute, parent gin.IRouter) {
	rt.setup(parent)
}

func (that *routeGroup) setup(parent gin.IRouter) {
	p := parent.Group(that.name)
	for _, rt := range that.routes {
		rt.setup(p)
	}
}

func (that *routeUse) setup(parent gin.IRouter) {
	parent.Use(that.middleware...)
	for _, rt := range that.routes {
		rt.setup(parent)
	}
}

func (that *route) setup(parent gin.IRouter) {
	typ := reflect.TypeOf(that.handlerFunc)

	where := fmt.Sprintf("route: %s, handlerFunc:%s", that.relativePath, typ.Name())
	if typ.Kind() != reflect.Func {
		panic("the route handlerFunc must be a function, " + where)
	}
	argNum := typ.NumIn()
	if argNum == 0 || argNum > 2 {
		panic("route handleFunc bad arguments, " + where)
	}

	if !typ.In(0).AssignableTo(typeGinContext) {
		panic("route handleFunc bad arguments, " + where)
	}
	if argNum == 1 {

		realHandleFunc, ok := that.handlerFunc.(func(ctx *gin.Context))
		if !ok {
			panic("type assertion fail, " + where)
		}
		parent.Handle(that.method, that.relativePath, realHandleFunc)

	} else if argNum == 2 {
		typeParam := typ.In(1)
		if typeParam.Kind() != reflect.Struct {
			panic("type assertion fail" + where)
		}
		_, ok := reflect.New(typeParam).Interface().(validate.ValidationAction)
		if !ok {
			panic("type assertion fail" + where)
		}
		f := reflect.ValueOf(that.handlerFunc)
		handleFunc := func(context *gin.Context) {
			param := reflect.New(typeParam).Interface().(validate.ValidationAction)
			if param.Validate(context, &param) {
				f.Call(valOf(context, &param))
			}
		}
		parent.Handle(that.method, that.relativePath, handleFunc)
	}
}

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}

func group(name string, r ...iRoute) *routeGroup {
	return &routeGroup{
		name:   name,
		routes: r,
	}
}

func use(middleware gin.HandlerFunc, r ...iRoute) *routeUse {
	return &routeUse{
		middleware: []gin.HandlerFunc{middleware},
		routes:     r,
	}
}

// TODO 2021-1-26 更新 handlerFunc 为准确的类型
func post(relativePath string, handlerFunc interface{}) *route {
	return handle(http.MethodPost, relativePath, handlerFunc)
}

func get(relativePath string, handlerFunc interface{}) *route {
	return handle(http.MethodGet, relativePath, handlerFunc)
}

func delete(relativePath string, handlerFunc interface{}) *route {
	return handle(http.MethodDelete, relativePath, handlerFunc)
}

func patch(relativePath string, handlerFunc interface{}) *route {
	return handle(http.MethodPatch, relativePath, handlerFunc)
}

func put(relativePath string, handlerFunc interface{}) *route {
	return handle(http.MethodPut, relativePath, handlerFunc)
}

func head(relativePath string, handlerFunc interface{}) *route {
	return handle(http.MethodHead, relativePath, handlerFunc)
}

func handle(method, relativePath string, handlerFunc interface{}) *route {
	return &route{
		method:       method,
		relativePath: relativePath,
		handlerFunc:  handlerFunc,
	}
}
