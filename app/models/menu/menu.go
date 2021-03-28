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
	Hidden      int    `json:"hidden"`
	Url         string `json:"url"`
	Icon        string `validate:"required" json:"icon"`
	Description string `validate:"required" json:"description"`
	Children    string `json:"children" gorm:"-"`
}

type MenuRouter struct {
	Id        int          `json:"id"`
	PId       int          `json:"p_id"`
	Name      string       `json:"name"`
	Url       string       `json:"url"`
	Hidden    int          `json:"hidden"`
	Component string       `json:"component"`
	Children  []MenuRouter `json:"children" gorm:"-"`
}

func (MenuModel) TableName() string {
	return "menus"
}
