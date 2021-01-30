package validates

import (
	"github.com/spf13/cast"
	"gocms/app/models/role"
	"gocms/app/validates/validate"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"gocms/pkg/response"
	"gocms/wrap"
	"net/http"
)

// 验证管理员创建参数
func VidateCreateOrUpdateRole(c *wrap.ContextWrapper, modelParams *role.RoleModel) bool {
	err := c.ShouldBindJSON(&modelParams)
	db := config.Db
	isExist := 0

	if err != nil {
		logger.PanicError(err, "参数验证 VidateCreateOrUpdateRole", false)
		response.ErrorResponse(http.StatusBadRequest, err.Error()).WriteTo(c)
		return false
	}

	if validate.WithResponseMsg(modelParams, c) == false {
		return false
	}

	if modelParams.ID > 0 {
		// 检查是否唯一
		db.Model(modelParams).Where("name = ? and id <> ?", modelParams.Name, modelParams.ID).Count(&isExist)
		if cast.ToBool(isExist) {
			response.ErrorResponse(http.StatusForbidden, "角色已存在").WriteTo(c)
			return false
		}
	} else {
		db.Model(modelParams).Where("name = ?", modelParams.Name).Count(&isExist)
		if cast.ToBool(isExist) {
			response.ErrorResponse(http.StatusForbidden, "角色已存在").WriteTo(c)
			return false
		}
	}
	return true
}
