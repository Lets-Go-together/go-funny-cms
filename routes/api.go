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
	menuController := new(Admin.MenuController)

	rt :=
		group("api",
			post("/login", authController.Login),
			post("/admin/register", authController.Register),
			//post("/admin/register2", authController.Register2),
			get("/pwd", toolController.Pwd),
			use(middleware.Authenticate,
				get("/me", authController.Me),
				delete_("/logout", authController.Logout),
				group("admin",
					get("", adminController.Index),
					get("/show", adminController.Show),
					post("/store", adminController.Store),
					put("/save", adminController.Save),
					delete_("/delete", adminController.Destory),
				),
				group("permission",
					get("", permissionController.Index),
					put("/save", permissionController.Save),
					get("/show", permissionController.Show),
					get("/tree", permissionController.Tree),
					post("/store", permissionController.Store),
					delete_("/delete", permissionController.Destory),
				),
				group("role",
					get("", roleController.Index),
					get("/show", roleController.Show),
					post("/store", roleController.Store),
					put("/save", roleController.Save),
					delete_("/delete", roleController.Destory),
				),
				group("menu",
					get("", menuController.Index),
					get("/show", menuController.Show),
					post("/store", menuController.Store),
					put("/save", menuController.Save),
					delete_("/delete", menuController.Destory),
					get("/tree", menuController.Tree),
				),
			),
			get("/qiniu", toolController.Qiniu),
		)
	rt.setup(engine)
}
