package middleware

import (
	"github.com/gin-gonic/gin"
	"gocms/pkg/logger"
	"net"
	"net/http"
	"os"
	"strings"
)

func DefaultMiddle(c *gin.Context) {

	c.Next()
}

func RecoveryMiddleware(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				if se, ok := ne.Err.(*os.SyscallError); ok {
					if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
						strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
						brokenPipe = true
					}
				}
			}
			logger.Error("Internal Server Error", err)
			if brokenPipe {
				_ = context.Error(err.(error))
				context.Abort()
			} else {
				context.AbortWithStatus(http.StatusInternalServerError)
			}
		}
	}()
	context.Next()
}
