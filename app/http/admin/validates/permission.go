package validates

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gocms/app/validates/validate"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"gocms/pkg/response"
	"net/http"
)

type Permission struct {
	Name   string `validate:"required" json:"name"`
	Icon   string `validate:"required" json:"name"`
	Url    string `validate:"required" json:"name"`
	Method string `validate:"required" json:"name"`
	Pid    int    `validate:"required,number" json:"pid"`
}

type PermissionUpdate struct {
	Permission
}

// 验证管理员创建参数
func VidateCreateOrUpdatePermission(c *gin.Context, params *map[string]string) bool {
	err := c.ShouldBindJSON(&params)
	db := config.Db
	modelParams := cast.ToStringMap(params)
	isExist := 0

	if err != nil {
		logger.PanicError(err, "参数验证 VidateCreateOrUpdatePermission", true)
	}

	if validate.WithResponseMsg(params, c) {
		return false
	}

	if _, id := modelParams["id"]; id == true {
		// 检查是否唯一
		db.Where("method = ? and url = ? and id <> ?", modelParams["method"], modelParams["url"], id).Count(&isExist)
		if cast.ToBool(isExist) == true {
			response.ErrorResponse(http.StatusForbidden, "权限已存在").WriteTo(c)
			return false
		}
	} else {
		db.Where("method = ? and url = ?", modelParams["method"], modelParams["url"]).Count(&isExist)
		if cast.ToBool(isExist) == true {
			response.ErrorResponse(http.StatusForbidden, "权限已存在").WriteTo(c)
			return false
		}
	}
	return true
}
