package routes

import (
	"github.com/gin-gonic/gin"
	Admin "gocms/app/http/admin/controllers"
	"gocms/app/http/middleware"
)

// 路由注册
func RegisterApiRoutes(engine *gin.Engine) {

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
				delete_("/logout", authController.Logout),
				group("admin",
					get("", adminController.Index),
					get("/:id", adminController.Show),
					post("", adminController.Store),
					put("/:id", adminController.Save),
					delete_("/:id", adminController.Destory),
				),
				group("permission",
					get("", permissionController.Index),
					get("/:id", permissionController.Show),
					post("", permissionController.Store),
					put("", permissionController.Save),
					delete_("/:id", permissionController.Destory),
				),
				group("role",
					get("", roleController.Index),
					get("/:id", roleController.Show),
					post("", roleController.Store),
					put("/:id", roleController.Save),
					delete_("/:id", roleController.Destory),
				),
			),
			get("/qiniu", toolController.Qiniu),
		)
	rt.setup(engine)
}
