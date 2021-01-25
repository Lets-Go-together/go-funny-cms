package routes

import (
	"github.com/gin-gonic/gin"
	Admin "gocms/app/http/admin/controllers"
	"gocms/app/http/middleware"
)

func route() {
	type route struct {
		method       string
		relativePath string
		handlerFunc  gin.HandlerFunc
		auth         bool
	}

	type group struct {
		name   string
		routes []*route
	}

	var post = func(relativePath string, handlerFunc gin.HandlerFunc) *route {
		return &route{
			method:       "",
			relativePath: "",
			handlerFunc:  handlerFunc,
			auth:         false,
		}
	}

	_ = group{
		name: "",
		routes: []*route{
			post("/user", nil),
		},
	}
}

// 路由注册
func RegisterApiRoutes(router *gin.Engine) {
	// 后台项目
	apiRouter := router.Group("api")

	// 不需要登陆的路由
	authController := new(Admin.AuthController)
	apiRouter.POST("/login", authController.Login)
	apiRouter.POST("/admin/register", authController.Register)

	ToolController := new(Admin.ToolController)
	apiRouter.GET("/pwd", ToolController.Pwd)

	// 需要鉴权的路由
	apiRouter.Use(middleware.Authenticate)
	{
		apiRouter.GET("/me", authController.Me)
		apiRouter.DELETE("/logout", authController.Logout)

		// resful api
		AdminController := new(Admin.AdminController)
		apiAdminRouter := apiRouter.Group("admin")
		{
			apiAdminRouter.GET("", AdminController.Index)
			apiAdminRouter.GET("/:id", AdminController.Show)
			apiAdminRouter.POST("", AdminController.Store)
			apiAdminRouter.PUT("/:id", AdminController.Save)
		}

		PermissionController := new(Admin.PermissionController)
		apiPermissionRouter := apiRouter.Group("permission")
		{
			apiPermissionRouter.GET("", PermissionController.Index)
			apiPermissionRouter.GET("/:id", PermissionController.Show)
			apiPermissionRouter.POST("", PermissionController.Store)
			apiPermissionRouter.PUT("/:id", PermissionController.Save)
			apiPermissionRouter.DELETE("/:id", PermissionController.Destory)
		}

		RoleController := new(Admin.RoleController)
		apiRoleRouter := apiRouter.Group("role")
		{
			apiRoleRouter.GET("", RoleController.Index)
			apiRoleRouter.GET("/:id", RoleController.Show)
			apiRoleRouter.POST("", RoleController.Store)
			apiRoleRouter.PUT("/:id", RoleController.Save)
			apiRoleRouter.DELETE("/:id", RoleController.Destory)
		}

		apiRouter.GET("/qiniu", ToolController.Qiniu)
	}
}

func interceptor() {

}
