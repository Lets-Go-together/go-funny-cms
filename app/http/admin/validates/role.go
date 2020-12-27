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

type Role struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
}

type RoleUpdate struct {
	Permission
}

// 验证管理员创建参数
func VidateCreateOrUpdateRole(c *gin.Context, params *map[string]string) bool {
	err := c.ShouldBindJSON(&params)
	db := config.Db
	modelParams := cast.ToStringMap(params)
	isExist := 0

	if err != nil {
		logger.PanicError(err, "参数验证 VidateCreateOrUpdateRole", true)
	}

	if validate.WithResponseMsg(params, c) {
		return false
	}

	if _, id := modelParams["id"]; id == true {
		// 检查是否唯一
		db.Where("name = ? and id <> ?", modelParams["name"], id).Count(&isExist)
		if cast.ToBool(isExist) == true {
			response.ErrorResponse(http.StatusForbidden, "角色已存在").WriteTo(c)
			return false
		}
	} else {
		db.Where("name = ?", modelParams["name"]).Count(&isExist)
		if cast.ToBool(isExist) == true {
			response.ErrorResponse(http.StatusForbidden, "角色已存在").WriteTo(c)
			return false
		}
	}
	return true
}
