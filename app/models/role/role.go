package role

import (
	"gocms/app/models/base"
)

type RoleModel struct {
	base.BaseModel
	Name            string              `validate:"required" json:"name"`
	Description     string              `validate:"required" json:"description"`
	Permissions     []map[string]string `json:"permissions" gorm:"-"`
	Permissions_ids []int               `json:"-" gorm:"-"`
}

func (RoleModel) TableName() string {
	return "roles"
}
