package service

import (
	"github.com/cheggaaa/pb/v3"
	"gocms/app/models/permission"
	"gocms/pkg/config"
	"gocms/pkg/help"
)

type PermissionService struct{}

var Enforcer = config.Enforcer

// 生成权限
// 用于生成系统和全局的权限列表
// 更新创建动作
func (*PermissionService) GeneratePermissionNodes() {
	routers := config.GetAllRoutes()
	db := config.Db
	count := len(routers)
	bar := pb.StartNew(count)
	for _, router := range routers {
		permissionModel := &permission.Permission{}
		routerCopy := &permission.Permission{
			Url:    router["url"],
			Method: router["method"],
		}
		db.Where(routerCopy).Select("id").First(permissionModel)
		if permissionModel.ID == 0 {
			routerCopy.Name = router["name"]
			db.Model(&permission.Permission{}).Omit("pid", "status", "icon", "hidden").Create(routerCopy)
		} else {
			db.Model(&permission.Permission{}).Where("id = ?", permissionModel.ID).Update(routerCopy)
		}
		bar.Increment()
	}
}

// 获取全部权限节点
// 参数为空的的时候：返回本地的路由中的权限节点
// 否则返回数据库的
func (*PermissionService) GetPermissionNodes(params ...interface{}) []map[string]string {
	if param := help.GetDefaultParam(params); param == nil {
		return config.GetAllRoutes()
	}
	return []map[string]string{}
}

// 检查是否有这个权限
// 参数必须存在
// 第三个参数可以为空，默认当前admin_id
func (*PermissionService) HasPermissionForNode(node string, method string, admin_id ...interface{}) {

}

// 授权
// 访问地址 string
// 动作	   string
// 账号    string
func (*PermissionService) GrantPermissionForAdmin(permission string, method string, account string) {

}
