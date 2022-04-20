package admin

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func Get(c *gin.Context) {
	zap.L().Info("user get function called.", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
	userId, _ := strconv.Atoi(c.Param("id"))
	// Get the user by the `username` from the database.
	user, err := model.GetUserByID(uint64(userId))
	// if user is not found or deleted, send error
	if err != nil {
		SendResponse(c, berror.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
