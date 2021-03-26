package service

import (
	"fmt"
	"gocms/app/models/base"
	"gocms/app/models/permission"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/wrap"
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
		if v.PId == pid {
			v.Children = getPermisstionTree(permissions, v.Id)
			list = append(list, v)
		}

	}

	return list
}

func (*PermissionService) GetList(page int, pageSize int, c *wrap.ContextWrapper) *base.Result {
	var permissions []PermissionList
	offset := help.GetOffset(page, pageSize)
	total := 0
	keyword := c.Query("keyword")

	query := config.Db.Model(&permission.Permission{}).Select("id, name, icon, url, status, method, p_id, hidden, created_at")

	if len(keyword) > 0 {
		query = query.Where("name|url like ?", fmt.Sprintf("%%s%", keyword))
	}

	query.Limit(pageSize).Offset(offset).Scan(&permissions)
	query.Count(&total)

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
	permissions = getPermisstionTree(permissions, 1)

	return permissions
}
