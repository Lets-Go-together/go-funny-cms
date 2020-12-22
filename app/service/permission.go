package service

import (
	"github.com/cheggaaa/pb/v3"
	"gocms/app/models"
	"gocms/pkg/config"
	"gocms/pkg/help"
)

type PermissionService struct{}

// 生成权限
// 用于生成系统和全局的权限列表
// 更新创建动作
func GeneratePermissionNodes() {
	routers := config.GetAllRoutes()
	db := config.Db
	count := len(routers)
	bar := pb.StartNew(count)
	for _, router := range routers {
		permissionModel := &models.Permission{}
		routerCopy := &models.Permission{
			Url:    router["url"],
			Method: router["method"],
		}
		db.Where(routerCopy).Select("id").First(permissionModel)
		if permissionModel.ID == 0 {
			routerCopy.Name = router["name"]
			db.Model(&models.Permission{}).Omit("pid", "status", "icon", "hidden").Create(routerCopy)
		} else {
			db.Model(&models.Permission{}).Where("id = ?", permissionModel.ID).Update(routerCopy)
		}
		bar.Increment()
	}
}

// 获取全部权限节点
// 参数为空的的时候：返回本地的路由中的权限节点
// 否则返回数据库的
func GetPermissionNodes(params ...interface{}) []map[string]string {
	if param := help.GetDefaultParam(params); param == nil {
		return config.GetAllRoutes()
	}
	return []map[string]string{}
}
