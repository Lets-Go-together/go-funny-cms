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
	Admin
	ID       uint64 `validate:"required"`
	Password string `validate:"gte=6,lte=16" json:"password"`
}

// 验证管理员创建参数
func VidateCreateAdmin(c *gin.Context, params *Admin) bool {
	err := c.ShouldBindJSON(&params)

	if err != nil {
		response.ErrorResponse(http.StatusForbidden, "请检查参数").WriteTo(c)
		return false
	}

	uniqueWheres := map[string]string{
		"email":   params.Email,
		"account": params.Email,
	}

	if r := isAllowCreateAdmin(uniqueWheres); r == false {
		response.ErrorResponse(http.StatusForbidden, "账号或者邮箱已存在").WriteTo(c)
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

	uniqueWheres := map[string]string{
		"email":   params.Email,
		"account": params.Email,
		"id":      cast.ToString(params.ID),
	}

	if r := isAllowCreateAdmin(uniqueWheres); r == false {
		response.ErrorResponse(http.StatusForbidden, "账号或者邮箱已存在").WriteTo(c)
		return false
	}

	return validate.WithResponseMsg(params, c)
}

// 批量验证是否可以创建
// == true 为是
func isAllowCreateAdmin(where map[string]string) bool {
	var total int
	db := config.Db.Model(&adminModel.Admin{})
	if _, ok := where["id"]; ok == true {
		delete(where, "id")
		db = db.Where("id", "<>", where["id"])
	}
	db.Where(where).Count(&total)

	return total == 0
}
