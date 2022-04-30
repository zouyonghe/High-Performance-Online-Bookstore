package common

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/handler/user"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// SelfDel deletes the user of token specified.
//
// @Summary SelfDel deletes the user of token specified
// @Description SelfDel deletes the user of token specified
// @Tags user/common
// @Produce  json
// @Success 200 {object} user.SwaggerSelfDelResponse "{"code":0,"message":"OK","data":{"UserID":8}}"
// @Router /user/common/ [delete]
// @Security ApiKeyAuth
func SelfDel(c *gin.Context) {
	log.SelfDelCalled(c)

	UserID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, berror.InternalServerError, nil)
	}
	if err := model.DeleteUser(UserID); err != nil {
		log.ErrDeleteUser(err)
		SendResponse(c, berror.ErrDeleteUser, nil)
		return
	}
	rsp := user.SelfDelResponse{
		UserID: UserID,
	}
	SendResponse(c, nil, rsp)
}
