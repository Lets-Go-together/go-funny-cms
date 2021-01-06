package validates

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gocms/app/models/admin"
	"gocms/app/service"
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

// 验证管理员创建参数
func VidateCreateOrUpdateAdmin(c *gin.Context, adminParams *admin.Admin) bool {
	err := c.ShouldBindJSON(&adminParams)

	logger.Info("admin", adminParams)

	if err != nil {
		response.ErrorResponse(http.StatusBadRequest, err.Error()).WriteTo(c)
		return false
	}

	if !validate.WithResponseMsg(adminParams, c) {
		return false
	}

	uniqueWheres := make(map[string]string)
	if adminParams.ID > 0 {
		uniqueWheres = map[string]string{
			"email":   adminParams.Account,
			"account": adminParams.Email,
			"id":      cast.ToString(adminParams.ID),
		}
	} else {
		uniqueWheres = map[string]string{
			"email":   adminParams.Email,
			"account": adminParams.Account,
		}
	}

	dbModel := config.Db.Model(&admin.Admin{})
	if r := service.IsAllowOperationModel(uniqueWheres, dbModel); r == false {
		response.ErrorResponse(http.StatusForbidden, "账号或者邮箱已存在").WriteTo(c)
		return false
	}

	return true
}
