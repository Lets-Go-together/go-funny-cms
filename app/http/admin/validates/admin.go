package validates

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
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

type AdminUpdate struct {
	Admin
	ID       uint64 `validate:"required"`
	Password string `validate:"gte=6,lte=16" json:"password"`
}

// 验证管理员创建参数
func VidateCreateOrUpdateAdmin(c *gin.Context, params *map[string]string) bool {
	err := c.ShouldBindJSON(&params)
	adminParams := cast.ToStringMap(params)
	uniqueWheres := make(map[string]string)
	var valudateErr error

	if _, ok := adminParams["id"]; ok == true {
		adminValidate := &AdminUpdate{}
		if valudateErr = mapstructure.Decode(adminParams, &adminValidate); valudateErr == nil {
			uniqueWheres = map[string]string{
				"email":   adminValidate.Email,
				"account": adminValidate.Email,
				"id":      cast.ToString(adminValidate.ID),
			}
		}
	} else {
		adminValidate := &Admin{}
		if valudateErr = mapstructure.Decode(adminParams, &adminValidate); valudateErr == nil {
			uniqueWheres = map[string]string{
				"email":   adminValidate.Email,
				"account": adminValidate.Email,
			}
		}
	}

	if valudateErr != nil {
		logger.PanicError(err, "validate", false)
		response.ErrorResponse(http.StatusForbidden, "系统错误").WriteTo(c)
		return false
	}

	dbModel := config.Db.Model(&admin.Admin{})
	if r := service.IsAllowOperationModel(uniqueWheres, dbModel); r == false {
		response.ErrorResponse(http.StatusForbidden, "账号或者邮箱已存在").WriteTo(c)
		return false
	}

	return validate.WithResponseMsg(params, c)
}
