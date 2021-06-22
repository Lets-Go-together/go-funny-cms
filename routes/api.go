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
	notificationController := new(Admin.NotificationController)
	MailController := new(Admin.MailController)
	schedulerController := new(Admin.SchedulerController)

	rt :=
		group("api",
			group("schedule",
				get("task", schedulerController.List),
				put("task", schedulerController.Add),
				delete_("task", schedulerController.Delete),
				patch("task", schedulerController.Update),
			),
			post("/login", authController.Login),
			get("/pwd", toolController.Pwd),
			use(middleware.Authenticate,
				get("/me", authController.Me),
				delete_("/logout", authController.Logout),
				use(middleware.Permission),
				group("admin",
					get("", adminController.Index),
					get("/show", adminController.Show),
					post("/store", adminController.Store),
					put("/save", adminController.Save),
					delete_("/delete", adminController.Destory),
					get("/tree", adminController.Tree),
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
				group("notify",
					get("", notificationController.List),
					post("read", notificationController.Readed),
					put("store", notificationController.Store),
				),
				group("mail",
					get("", MailController.List),
					get("mailer", MailController.Mailer),
					post("store", MailController.Store),
					post("resend", MailController.Resend),
					delete_("delete", MailController.Delete),
				),
			),
			get("/qiniu", toolController.Qiniu),
		)
	rt.setup(engine)
}
