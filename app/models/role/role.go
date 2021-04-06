package role

import (
	"gocms/app/models/base"
)

type RoleModel struct {
	base.BaseModel
	Name            string              `json:"name"`
	Status          int                 `json:"status"`
	Description     string              `json:"description"`
	MenuIds         base.IntJson        `json:"menu_ids"`
	Permissions     []map[string]string `json:"permissions" gorm:"-"`
	Permissions_ids []int               `json:"permission_ids" gorm:"-"`
}

func (RoleModel) TableName() string {
	return "roles"
}
