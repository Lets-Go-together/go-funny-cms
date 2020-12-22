package controllers

import (
	"github.com/gin-gonic/gin"
	"gocms/app/service"
)

type PermissionController struct{}

var permissionService = &service.PermissionService{}

// 权限节点列表
func (that *PermissionController) Index(c *gin.Context) {

}
