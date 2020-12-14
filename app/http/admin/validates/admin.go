package validates

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	adminModel "gocms/app/models/admin"
	"gocms/app/validates/validate"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"gocms/pkg/response"
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

type AdminUpdate struct {
	ID          uint64 `validate:"required"`
	Description string `validate:"required,gte=1,lte=60" json:"description"`
	Email       string `validate:"required,email" json:"email"`
	Account     string `validate:"required,gte=2,lte=8" json:"account"`
	Password    string `validate:"gte=6,lte=16" json:"password"`
	Phone       string `validate:"required" json:"phone"`
	Avatar      string `validate:"required,url" json:"avatar"`
}

// 验证管理员创建参数
func VidateCreateAdmin(c *gin.Context, params *Admin) bool {
	err := c.ShouldBindJSON(&params)

	if err != nil {
		response.ErrorResponse(http.StatusForbidden, "请检查参数").WriteTo(c)
		return false
	}
	var admin = adminModel.Admin{}
	config.Db.Select("account").Where(adminModel.Admin{
		Account: params.Account,
	}).First(&admin)

	if len(admin.Account) > 0 {
		response.ErrorResponse(http.StatusForbidden, "账号已存在").WriteTo(c)
		return false
	}

	return validate.WithResponseMsg(params, c)
}

// 验证管理员更新参数
func VidateUpdateAdmin(c *gin.Context, params *AdminUpdate) bool {
	err := c.ShouldBindJSON(&params)
	info, _ := json.Marshal(params)
	logger.Info("params", string(info))

	if err != nil {
		response.ErrorResponse(http.StatusForbidden, "请检查参数").WriteTo(c)
		return false
	}
	var admin = adminModel.Admin{}
	config.Db.Select("account").Where("id <> " + cast.ToString(params.ID)).Where(adminModel.Admin{
		Account: params.Account,
	}).First(&admin)

	if len(admin.Account) > 0 {
		response.ErrorResponse(http.StatusForbidden, "账号已存在").WriteTo(c)
		return false
	}

	return validate.WithResponseMsg(params, c)
}
