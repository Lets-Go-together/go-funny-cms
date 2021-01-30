package validates

import (
	"github.com/spf13/cast"
	"gocms/app/models/permission"
	"gocms/app/validates/validate"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"gocms/pkg/response"
	"gocms/wrap"
	"net/http"
)

// 验证管理员创建参数
func VidateCreateOrUpdatePermission(c *wrap.ContextWrapper, modelParams *permission.Permission) bool {
	err := c.ShouldBindJSON(&modelParams)
	db := config.Db
	isExist := 0

	if err != nil {
		logger.PanicError(err, "参数验证 VidateCreateOrUpdatePermission", false)
		response.ErrorResponse(http.StatusBadRequest, err.Error()).WriteTo(c)
		return false
	}

	if !validate.WithResponseMsg(modelParams, c) {
		return false
	}

	if modelParams.ID > 0 {
		// 检查是否唯一
		db.Model(modelParams).Where("method = ? and url = ? and id <> ?", modelParams.Method, modelParams.Url, modelParams.ID).Count(&isExist)
		if cast.ToBool(isExist) {
			response.ErrorResponse(http.StatusForbidden, "权限已存在").WriteTo(c)
			return false
		}
	} else {
		db.Model(modelParams).Where("method = ? and url = ?", modelParams.Method, modelParams.Url).Count(&isExist)
		if cast.ToBool(isExist) {
			response.ErrorResponse(http.StatusForbidden, "权限已存在").WriteTo(c)
			return false
		}
	}

	return true
}
