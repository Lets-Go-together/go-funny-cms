package permission

import (
	"github.com/jinzhu/gorm"
	"gocms/app/models/base"
	"gocms/pkg/config"
)

type Permission struct {
	base.BaseModel
	Name   string `validate:"required" json:"name" binding:"required"`
	Icon   string `json:"icon" gorm:"-"`
	Url    string `validate:"required" json:"url"`
	Status int    `json:"status" gorm:"-"`
	Method string `validate:"required" json:"method"`
	PId    int    `validate:"required" json:"p_id"`
}

func (Permission) TableName() string {
	return "permissions"
}

func Instance() *gorm.DB {
	return config.Db.Model(&Permission{})
}
