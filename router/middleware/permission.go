package middleware

import (
	"Jinshuzhai-Bookstore/pkg/token"
	"github.com/gin-gonic/gin"
)

func HasPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, _ := token.ParseRequest(c)
		if ctx.Role != "admin" {
			c.JSON(200, gin.H{
				"code":    403,
				"message": "Permission denied",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
