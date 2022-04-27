package common

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/log"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/service"
	"github.com/gin-gonic/gin"
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
	log.SelfDelCalled(c)

	userId, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, berror.InternalServerError, nil)
	}
	if err := model.DeleteUser(userId); err != nil {
		log.ErrDeleteUser(err)
		SendResponse(c, berror.ErrDeleteUser, nil)
		return
	}
	rsp := user.SelfDelResponse{
		UserId: userId,
	}
	SendResponse(c, nil, rsp)
}
