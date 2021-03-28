package menu

import (
	"gocms/app/models/base"
)

type MenuModel struct {
	base.BaseModel
	Name        string `validate:"required" json:"name"`
	PId         int    `validate:"required" json:"p_id"`
	Weight      int    `json:"weight"`
	Component   string `json:"component"`
	Status      int    `validate:"required" json:"status"`
	Icon        string `validate:"required" json:"icon"`
	Description string `validate:"required" json:"description"`
	Children    string `json:"children" gorm:"-"`
}

type MenuRouter struct {
	Id        int          `json:"id"`
	PId       int          `json:"p_id"`
	Name      string       `json:"name"`
	Component string       `json:"component"`
	Children  []MenuRouter `json:"children" gorm:"-"`
}

func (MenuModel) TableName() string {
	return "menus"
}
