package middleware

import (
	"Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, berror.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
