package controllers

import (
	"github.com/gin-gonic/gin"
	authValidate "gocms/app/validates/auth"
	"gocms/pkg/response"
	"net/http"
)

type AuthController struct{}

func (*AuthController) Login(c *gin.Context) {
	authValidateAction := authValidate.LoginAction{}
	if msg := authValidateAction.Validate(c); len(msg) > 0 {
		c.JSON(http.StatusOK, response.JsonResponse{
			Status:  403,
			Message: msg,
		})
		return
	}

	//params := authValidateAction.GetLoginData()

	c.JSON(http.StatusOK, response.JsonResponse{
		Status:  200,
		Message: "Success",
	})
}

func (*AuthController) Register(c *gin.Context) {

}
