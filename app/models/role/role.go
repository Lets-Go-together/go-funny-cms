package role

import (
	"gocms/app/models/base"
)

type RoleModel struct {
	base.BaseModel
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions" gorm:"-"`
}

func (RoleModel) TableName() string {
	return "roles"
}
