package middleware

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
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
	sub := ctx.Role
	obj := c.Request.URL.Path
	act := c.Request.Method
	log.CheckPermissionCalled(c, sub, obj, act)

	if permission.CheckPermission(sub, obj, act) {
		c.Next()
		return
	} else {
		SendDenyResponse(c)
		c.Abort()
		return
	}
}
