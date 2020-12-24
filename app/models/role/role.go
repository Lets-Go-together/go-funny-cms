package role

import (
	"gocms/app/models/base"
)

type RoleModel struct {
	base.BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (RoleModel) TableName() string {
	return "roles"
}
