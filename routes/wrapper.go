package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gocms/app/validates/validate"
	"gocms/wrap"
	"net/http"
	"reflect"
)

var typeGinContext = reflect.TypeOf((*wrap.ContextWrapper)(nil))

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

		realHandleFunc, ok := that.handlerFunc.(func(ctx *wrap.ContextWrapper))
		if !ok {
			panic("type assertion fail, " + where)
		}
		parent.Handle(that.method, that.relativePath, func(context *gin.Context) {
			realHandleFunc(wrap.Context(context))
		})

	} else if argNum == 2 {

		typeParam := typ.In(1).Elem()
		if typeParam.Kind() != reflect.Struct {
			panic("type assertion fail" + where)
		}
		_, ok := reflect.New(typeParam).Interface().(validate.Validatable)
		if !ok {
			panic("type assertion fail" + where)
		}
		valueFunc := reflect.ValueOf(that.handlerFunc)

		proxyHandleFunc := func(context *gin.Context) {
			wrapCtx := wrap.Context(context)
			// 调用真实的 Router HandlerFunc
			invokeRealHandler(wrapCtx, typeParam, valueFunc)
		}
		parent.Handle(that.method, that.relativePath, proxyHandleFunc)
	}
}

func invokeRealHandler(context *wrap.ContextWrapper, tParam reflect.Type, vRealHandlerFunc reflect.Value) {
	param := reflect.New(tParam).Interface().(validate.Validatable)
	if param.Validate(context) {
		vRealHandlerFunc.Call(valOf(context, reflect.ValueOf(param).Interface()))
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

func delete_(relativePath string, handlerFunc interface{}) *route {
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
