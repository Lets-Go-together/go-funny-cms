package permission

import (
	"gocms/app/models/base"
)

type Permission struct {
	base.BaseModel
	Name   string `validate:"required" json:"name" binding:"required"`
	Icon   string `validate:"required" json:"icon" gorm:"-"`
	Url    string `validate:"required" json:"url"`
	Status int    `json:"status" gorm:"-"`
	Method string `validate:"required" json:"method"`
	Pid    int    `validate:"required" json:"pid" gorm:"-"`
}

func (Permission) TableName() string {
	return "permissions"
}
