package middleware

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/permission"
	"High-Performance-Online-Bookstore/pkg/token"
	"github.com/gin-gonic/gin"
)

// HasPermission checks if the user has the permission
// to access the resource.
func HasPermission(c *gin.Context) {
	var ctx *token.Context
	ctx, err := token.ParseRequest(c)
	if err != nil {
		ctx.Role = "guest"
	}
	if permission.CheckPermission(c, ctx.Role, c.Request.URL.Path, c.Request.Method) {
		c.Next()
		return
	} else {
		SendDenyResponse(c)
		c.Abort()
		return
	}
}
