package controllers

import (
	"github.com/gin-gonic/gin"
	"gocms/app/service"
	"gocms/pkg/response"
)

type AdminController struct{}

// 管理员列表
func (*AdminController) Index(c *gin.Context) {
	adminService := &service.AdminService{}
	list := adminService.GetList()
	response.SuccessResponse(list).WriteTo(c)
}
