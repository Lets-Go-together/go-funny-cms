package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller interface {
}

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
	handlerFunc  gin.HandlerFunc
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
	if rtGroup, ok := rt.(routeGroup); ok {
		rtGroup.setup(parent)
	}
	if rtRoute, ok := rt.(route); ok {
		rtRoute.setup(parent)
	}
}

func (that routeGroup) setup(parent gin.IRouter) {
	p := parent.Group(that.name)
	for _, rt := range that.routes {
		setupRoutes(rt, p)
	}
}

func (that routeUse) setup(parent gin.IRouter) {
	//p := parent.Use(that.middleware...)
	//for _, rt := range that.routes {
	//	setupRoutes(rt, p)
	//}
}

func (that route) setup(parent gin.IRouter) {
	parent.Handle(that.method, that.relativePath, that.handlerFunc)
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

func post(relativePath string, handlerFunc gin.HandlerFunc) *route {
	return handle(http.MethodPost, relativePath, handlerFunc)
}

func get(relativePath string, handlerFunc gin.HandlerFunc) *route {
	return handle(http.MethodGet, relativePath, handlerFunc)
}

func delete(relativePath string, handlerFunc gin.HandlerFunc) *route {
	return handle(http.MethodDelete, relativePath, handlerFunc)
}

func patch(relativePath string, handlerFunc gin.HandlerFunc) *route {
	return handle(http.MethodPatch, relativePath, handlerFunc)
}

func put(relativePath string, handlerFunc gin.HandlerFunc) *route {
	return handle(http.MethodPut, relativePath, handlerFunc)
}

func head(relativePath string, handlerFunc gin.HandlerFunc) *route {
	return handle(http.MethodHead, relativePath, handlerFunc)
}

func handle(method, relativePath string, handlerFunc gin.HandlerFunc) *route {
	return &route{
		method:       method,
		relativePath: relativePath,
		handlerFunc:  handlerFunc,
	}
}
