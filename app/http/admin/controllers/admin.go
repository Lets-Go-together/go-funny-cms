package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gocms/app/service"
	"gocms/pkg/response"
)

type AdminController struct{}

// 管理员列表
func (*AdminController) Index(c *gin.Context) {
	adminService := &service.AdminService{}
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "20")
	list := adminService.GetList(cast.ToInt(page), cast.ToInt(pageSize))

	response.SuccessResponse(list).WriteTo(c)
}

// 管理员创建
func (*AdminController) Create(c *gin.Context) {

}
