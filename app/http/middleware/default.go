package middleware

import "github.com/gin-gonic/gin"

func DefaultMiddle(c *gin.Context) {

	c.Next()
}
