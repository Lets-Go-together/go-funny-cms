package service

import (
	"fmt"
	"gocms/app/models"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
)

type AdminService struct{}
type listStruct struct {
	ID          uint64        `json:"id"`
	Account     string        `json:"account"`
	Description string        `json:"description"`
	Email       string        `json:"email"`
	Phone       string        `json:"phone"`
	Avatar      string        `json:"avatar"`
	CreatedAt   models.TimeAt `json:"created_at"`
}

func (*AdminService) GetList(page int, pageSize int) *models.Result {
	admins := []listStruct{}
	offset := help.GetOffset(page, pageSize)
	total := 0

	config.Db.Model(&models.Admin{}).Select("id, account, description, email, avatar, phone, created_at").Limit(pageSize).Offset(offset).Scan(&admins)
	config.Db.Model(&models.Admin{}).Count(&total)

	data := models.Result{
		Page:     page,
		PageSize: pageSize,
		List:     admins,
		Total:    total,
	}

	return &data
}

// 创建一个admin用户
func (*AdminService) Create(admin models.Admin) bool {
	r := config.Db.Omit("updated_at", "created_at").Create(&admin)
	if errs := r.GetErrors(); len(errs) > 0 {
		fmt.Println(errs[0])
		return false
	}

	return true
}

// 更新一个admin用户
func (*AdminService) Update(admin models.Admin, id string) bool {
	r := config.Db.Omit("updated_at", "created_at").Where("id = ?", id).Update(&admin)
	if errs := r.GetErrors(); len(errs) > 0 {
		logger.PanicError(errs[0], "更新admin用户", false)
		return false
	}

	return true
}
