package controllers

import (
	"github.com/gin-gonic/gin"
	"gocms/app/models/admin"
	authValidate "gocms/app/validates/auth"
	"gocms/pkg/config"
	"gocms/pkg/response"
	"net/http"
)

type AuthController struct{}

var Db = config.Db

func (*AuthController) Login(c *gin.Context) {
	authValidateAction := authValidate.LoginAction{}
	if msg := authValidateAction.Validate(c); len(msg) > 0 {
		c.JSON(http.StatusOK, response.JsonResponse{
			Status:  403,
			Message: msg,
		})
		return
	}

	params := authValidateAction.GetLoginData()
	adminModel := admin.Admin{}
	config.Db.Where("account = ?", params.Account).Where("password = ?", params.Password).Find(&adminModel)
	if adminModel.ID == 0 {
		c.JSON(http.StatusOK, response.JsonResponse{
			Status:  403,
			Message: "用户不存在",
			Data:    adminModel,
		})

		return
	}

	c.JSON(http.StatusOK, response.JsonResponse{
		Status:  200,
		Message: "Success",
		Data:    adminModel,
	})
}

func (*AuthController) Register(c *gin.Context) {

}
