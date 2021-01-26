package routes

import (
	"github.com/gin-gonic/gin"
	Admin "gocms/app/http/admin/controllers"
	"gocms/app/http/middleware"
)

// 路由注册
func RegisterApiRoutes2(router *gin.Engine) {

	authController := new(Admin.AuthController)
	toolController := new(Admin.ToolController)
	adminController := new(Admin.AdminController)
	permissionController := new(Admin.PermissionController)
	roleController := new(Admin.RoleController)

	rt :=
		group("api",
			post("/login", authController.Login),
			post("/admin/register", authController.Register),
			post("/admin/register2", authController.Register2),
			get("/pwd", toolController.Pwd),
			use(middleware.Authenticate,
				get("/me", authController.Me),
				delete("/logout", authController.Logout),
				group("admin",
					get("", adminController.Index),
					get("/:id", adminController.Show),
					post("", adminController.Store),
					put("/:id", adminController.Save),
				),
				group("permission",
					get("", permissionController.Index),
					get("/:id", permissionController.Show),
					post("", permissionController.Store),
					put("/:id", permissionController.Save),
					delete("/:id", permissionController.Destory),
				),
				group("role",
					get("", roleController.Index),
					get("/:id", roleController.Show),
					post("", roleController.Store),
					put("/:id", roleController.Save),
					delete("/:id", roleController.Destory),
				),
			),
			get("/qiniu", toolController.Qiniu),
		)
	setupRoutes(rt, router)
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
