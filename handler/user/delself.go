package user

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/pkg/token"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Delself deletes the user of token specified
func Delself(c *gin.Context) {
	ctx, _ := token.ParseRequest(c)
	id := ctx.ID
	if err := model.DeleteUser(id); err != nil {
		zap.L().Error("Delete self user failed", zap.Error(err))
		SendResponse(c, berror.ErrDeleteUser, nil)
		return
	}
	SendResponse(c, nil, gin.H{
		"ID":     id,
		"status": "deleted",
	})
}
