package validates

import (
	"github.com/jinzhu/gorm"
	"gocms/app/models/admin"
	"gocms/app/validates/validate"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"gocms/pkg/response"
	"gocms/wrap"
	"net/http"
)

type Admin struct {
	Description string `validate:"required,gte=1,lte=60" json:"description"`
	Email       string `validate:"required,email" json:"email"`
	Account     string `validate:"required,gte=2,lte=8" json:"account"`
	Password    string `validate:"required,gte=6,lte=16" json:"password"`
	Phone       string `validate:"required" json:"phone"`
	Avatar      string `validate:"required,url" json:"avatar"`
}

// 验证管理员创建参数
func VidateCreateOrUpdateAdmin(c *wrap.ContextWrapper, adminParams *admin.Admin) bool {
	err := c.ShouldBindJSON(&adminParams)

	logger.Info("admin", adminParams)

	if err != nil {
		response.ErrorResponse(http.StatusBadRequest, err.Error()).WriteTo(c)
		return false
	}

	if !validate.WithResponseMsg(adminParams, c) {
		return false
	}

	modelQuery := func() *gorm.DB {
		return config.Db.Model(admin.Admin{})
	}

	if adminParams.ID > 0 {
		total := 0
		modelQuery().Where("id <> ?", adminParams.ID).Where("account = ?", adminParams.Account).Count(&total)
		if total > 0 {
			response.ErrorResponse(http.StatusForbidden, "账号已存在").WriteTo(c)
			return false
		}

		modelQuery().Where("id <> ?", adminParams.ID).Where("email = ?", adminParams.Email).Count(&total)
		if total > 0 {
			response.ErrorResponse(http.StatusForbidden, "邮箱已存在").WriteTo(c)
			return false
		}
	} else {
		total := 0
		modelQuery().Where("account = ?", adminParams.Account).Count(&total)
		if total > 0 {
			response.ErrorResponse(http.StatusForbidden, "账号已存在").WriteTo(c)
			return false
		}

		modelQuery().Where("email = ?", adminParams.Email).Count(&total)
		if total > 0 {
			response.ErrorResponse(http.StatusForbidden, "邮箱已存在").WriteTo(c)
			return false
		}
	}

	return true
}
