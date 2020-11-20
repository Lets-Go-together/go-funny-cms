package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (*IndexController) Index(c *gin.Context) {
	fmt.Println("Hello Word")
}
