package middleware

import "github.com/gin-gonic/gin"

func Authenticate(c *gin.Context) {

	c.Next()
}
