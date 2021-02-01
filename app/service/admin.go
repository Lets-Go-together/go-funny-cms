package service

import (
	"fmt"
	"gocms/app/models/admin"
	"gocms/app/models/base"
	"gocms/pkg/auth"
	"gocms/pkg/auth/rabc"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
	"gocms/wrap"
)

type AdminService struct{}
type listStruct struct {
	ID          uint64      `json:"id"`
	Account     string      `json:"account"`
	Description string      `json:"description"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	Avatar      string      `json:"avatar"`
	CreatedAt   base.TimeAt `json:"created_at"`
	UpdatedAt   base.TimeAt `json:"updated_at"`
}

func (*AdminService) GetList(page int, pageSize int, c *wrap.ContextWrapper) *base.Result {
	admins := []listStruct{}
	offset := help.GetOffset(page, pageSize)
	total := 0
	account := c.Query("account")

	query := config.Db.Model(&admin.Admin{}).Select("id, account, description, email, avatar, phone, updated_at")
	if len(account) > 0 {
		query = query.Where("account like ?", "%"+account+"%")
	}

	query.Limit(pageSize).Offset(offset).Scan(&admins)
	config.Db.Model(&admin.Admin{}).Count(&total)

	data := base.Result{
		Page:     page,
		PageSize: pageSize,
		List:     admins,
		Total:    total,
	}

	return &data
}

// 创建一个admin用户
func (*AdminService) Create(admin admin.Admin) bool {
	r := config.Db.Omit("updated_at", "created_at").Create(&admin)
	if errs := r.GetErrors(); len(errs) > 0 {
		return false
	}

	return true
}

// 更新一个admin用户
func (*AdminService) Update(admin admin.Admin, id string) bool {
	r := config.Db.Omit("updated_at", "created_at").Where("id = ?", id).Update(&admin)
	if errs := r.GetErrors(); len(errs) > 0 {
		logger.PanicError(errs[0], "更新admin用户", false)
		return false
	}

	return true
}

// UpdateOrCreate 创建或者更新权限
func (*AdminService) UpdateOrCreate(adminModel admin.Admin, c *wrap.ContextWrapper) bool {

	if len(adminModel.Password) > 0 {
		adminModel.Password = auth.CreatePassword(adminModel.Password)
	}

	if adminModel.ID > 0 {
		config.Db.Model(adminModel).Update(&adminModel)
	} else {
		config.Db.Model(adminModel).Create(&adminModel)
	}

	fmt.Println(adminModel)

	rabc.DeleteRolesForUser(adminModel.Account)
	if len(adminModel.Role_ids) > 0 {
		roleNames := GetRolesName(adminModel.Role_ids)
		logger.Info("roles", roleNames)
		rabc.AddRolesForUser(adminModel.Account, roleNames)
	}

	return true
}
