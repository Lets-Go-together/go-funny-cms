package permission

import (
	"gocms/app/models/base"
)

type Permission struct {
	base.BaseModel
	Name   string `json:"name" binding:"required"`
	Icon   string `json:"icon" gorm:"-"`
	Url    string `json:"url"`
	Status int    `json:"status" gorm:"-"`
	Method string `json:"method"`
	Pid    int    `json:"pid" gorm:"-"`
}

func (Permission) TableName() string {
	return "permissions"
}
