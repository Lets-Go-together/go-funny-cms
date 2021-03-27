package role

import (
	"gocms/app/models/base"
)

type MenuModel struct {
	base.BaseModel
	Name        string `validate:"required" json:"name"`
	PId         int    `validate:"required" json:"p_id"`
	Status      int    `validate:"required" json:"status"`
	Icon        string `validate:"required" json:"icon"`
	Description string `validate:"required" json:"description"`
	Children    string `json:"children" gorm:"-"`
}

func (MenuModel) TableName() string {
	return "menus"
}
