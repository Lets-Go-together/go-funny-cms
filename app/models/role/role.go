package role

import (
	"gocms/app/models/base"
)

type RoleModel struct {
	base.BaseModel
	Name        string        `validate:"required" json:"name"`
	Description string        `validate:"required" json:"description"`
	Permissions []interface{} `json:"permissions" gorm:"-"`
}

func (RoleModel) TableName() string {
	return "roles"
}
