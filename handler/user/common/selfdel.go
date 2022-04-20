package common

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/pkg/token"
	"Jinshuzhai-Bookstore/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SelfDel deletes the user of token specified.
//
// @Summary SelfDel deletes the user of token specified
// @Description SelfDel deletes the user of token specified
// @Tags user/common
// @Produce  json
// @Success 200 {object} user.SwaggerSelfDelResponse "{"code":0,"message":"OK","data":{"userId":8}}"
// @Router /user/common/ [delete]
// @Security ApiKeyAuth
func SelfDel(c *gin.Context) {
	zap.L().Info("delete self function called", zap.String("X-Request-Id", util.GetReqID(c)))
	ctx, _ := token.ParseRequest(c)
	userId := ctx.ID
	if err := model.DeleteUser(userId); err != nil {
		zap.L().Error("Delete self user failed", zap.Error(err))
		SendResponse(c, berror.ErrDeleteUser, nil)
		return
	}
	rsp := user.SelfDelResponse{
		UserId: userId,
	}
	SendResponse(c, nil, rsp)
}
