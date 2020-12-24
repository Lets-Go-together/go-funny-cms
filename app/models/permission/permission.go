package permission

import (
	"gocms/app/models/base"
)

type Permission struct {
	base.BaseModel
	Name   string `json:"account"`
	Icon   string `json:"password,omitempty"`
	Url    string `json:"description"`
	Status int    `json:"email,omitempty"`
	Method string `json:"phone"`
	Pid    int    `json:"avatar,omitempty"`
}

func (Permission) TableName() string {
	return "permissions"
}
