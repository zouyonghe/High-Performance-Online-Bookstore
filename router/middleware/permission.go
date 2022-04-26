package middleware

import (
	"Jinshuzhai-Bookstore/pkg/token"
	"Jinshuzhai-Bookstore/policy"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// HasPermission checks if the user has the permission
// to access the resource.
func HasPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, err := token.ParseRequest(c)
		if err != nil {
			zap.L().Error("failed to parse request, default use guest role", zap.Error(err))
			ctx.Role = "guest"
		}
		if policy.CheckPermission(c, ctx.Role, c.Request.URL.Path, c.Request.Method) {
			c.Next()
			return
		} else {
			c.JSON(403, gin.H{
				"code": 403,
				"msg":  "Forbidden",
			})
			c.Abort()
			return
		}
	}
}
