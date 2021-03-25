package service

import (
	"gocms/app/models/base"
	"gocms/app/models/permission"
	"gocms/pkg/config"
	"gocms/pkg/help"
)

type PermissionService struct{}
type PermissionList struct {
	Id        int              `json:"id"`
	Name      string           `json:"name"`
	Icon      string           `json:"icon"`
	Url       string           `json:"url"`
	Status    int              `json:"status"`
	Hidden    int              `json:"hidden"`
	Method    string           `json:"method"`
	PId       int              `json:"p_id"`
	CreatedAt base.TimeAt      `json:"created_at"`
	Children  []PermissionList `json:"children"`
}

// https://binglangnet.com/vue/#/article-details/18?ii=3
// https://www.kancloud.cn/hyckrrwzsn/go/512829
func getPermisstionTree(permissions []PermissionList, pid int) []PermissionList {
	var list []PermissionList
	for _, v := range permissions {
		if v.Id == pid {
			v.Children = getPermisstionTree(permissions, v.Id)
		}
		list = append(list, v)
	}

	return list
}

func (*PermissionService) GetList(page int, pageSize int) *base.Result {
	var permissions []PermissionList
	offset := help.GetOffset(page, pageSize)
	total := 0

	config.Db.Model(&permission.Permission{}).Select("id, name, icon, url, status, method, p_id, hidden, created_at").Limit(pageSize).Offset(offset).Scan(&permissions)
	config.Db.Model(&permission.Permission{}).Count(&total)

	permissions = getPermisstionTree(permissions, 0)
	data := base.Result{
		Page:     page,
		PageSize: pageSize,
		List:     permissions,
		Total:    total,
	}

	return &data
}

// UpdateOrCreate 创建或者更新权限
func (*PermissionService) UpdateOrCreate(permissionModel permission.Permission) bool {
	if permissionModel.ID > 0 {
		return config.Db.Model(permissionModel).Update(permissionModel).RowsAffected > 0
	}

	return config.Db.Model(permissionModel).Create(permissionModel).RowsAffected > 0
}

// 获取权限节点树
func (*PermissionService) GetPermisstionTree() []PermissionList {
	var permissions []PermissionList
	config.Db.Model(&permission.Permission{}).Select("id, name, icon, url, status, method, p_id, hidden, created_at").Scan(&permissions)

	return getPermisstionTree(permissions, 0)
}
