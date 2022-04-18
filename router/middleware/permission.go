package middleware

import (
	"Jinshuzhai-Bookstore/pkg/token"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// HasPermission checks if the user has the permission
// to access the resource.
func HasPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, _ := token.ParseRequest(c)
		if ctx.Role != "admin" {
			c.JSON(200, gin.H{
				"code":    403,
				"message": "Permission denied",
			})
			c.Abort()
			zap.L().Info("Not admin user")
			return
		}
		c.Next()
	}
}
