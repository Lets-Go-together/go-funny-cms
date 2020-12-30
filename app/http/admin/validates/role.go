package validates

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gocms/app/models/role"
	"gocms/app/validates/validate"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"gocms/pkg/response"
	"net/http"
)

type Role struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
}

type RoleUpdate struct {
	Permission
}

// 验证管理员创建参数
func VidateCreateOrUpdateRole(c *gin.Context, modelParams *role.RoleModel) bool {
	err := c.ShouldBindJSON(&modelParams)
	db := config.Db
	isExist := 0

	if err != nil {
		logger.PanicError(err, "参数验证 VidateCreateOrUpdateRole", true)
	}

	if validate.WithResponseMsg(modelParams, c) {
		return false
	}

	if modelParams.ID > 0 {
		// 检查是否唯一
		db.Where("name = ? and id <> ?", modelParams.Name, modelParams.ID).Count(&isExist)
		if cast.ToBool(isExist) == true {
			response.ErrorResponse(http.StatusForbidden, "角色已存在").WriteTo(c)
			return false
		}
	} else {
		db.Where("name = ?", modelParams.Name).Count(&isExist)
		if cast.ToBool(isExist) == true {
			response.ErrorResponse(http.StatusForbidden, "角色已存在").WriteTo(c)
			return false
		}
	}
	return true
}
