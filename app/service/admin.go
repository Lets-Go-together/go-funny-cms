package service

import (
	"gocms/app/models/admin"
	"gocms/app/models/model"
	"gocms/pkg/config"
	"gocms/pkg/help"
)

type AdminService struct{}
type listStruct struct {
	ID          uint64       `json:"id"`
	Account     string       `json:"account"`
	Description string       `json:"description"`
	Email       string       `json:"email"`
	Phone       string       `json:"phone"`
	Avatar      string       `json:"avatar"`
	CreatedAt   model.TimeAt `json:"created_at"`
}

func (*AdminService) GetList(page int, pageSize int) *model.Result {
	admins := []listStruct{}
	offset := help.GetOffset(page, pageSize)
	total := 0

	config.Db.Model(&admin.Admin{}).Select("id, account, description, email, avatar, phone, created_at").Limit(pageSize).Offset(offset).Scan(&admins)
	config.Db.Model(&admin.Admin{}).Count(&total)

	data := model.Result{
		Page:     page,
		PageSize: pageSize,
		List:     admins,
		Total:    total,
	}

	return &data
}
