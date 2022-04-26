package admin

import (
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/service"
	"Jinshuzhai-Bookstore/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
import . "Jinshuzhai-Bookstore/handler"

// ListUser lists users account by specified username format.
//
// @Summary List users account by specified username format.
// @Description List users account by specified username format include id, username, encrypted password, etc.
// @Tags user/admin
// @Produce  json
// @Success 200 {object} user.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":1,"username":"admin","ShortId":"5P9Ia4QnR","password":"$2a$10$Fv9BWzqsiQ.JuuGdcXdvN.Fx3ml.dVR47W22GoJMWQAlm9wHQIMVe","role":"admin","createdAt":"2021-04-18 15:40:33","updatedAt":"2021-04-18 15:40:33"}]}}"
// @Router /user/admin [get]
// @Security ApiKeyAuth
func ListUser(c *gin.Context) {
	zap.L().Info("List user function called", zap.String("X-Request-Id", util.GetReqID(c)))
	var r user.ListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		zap.L().Error("Bind error", zap.Error(err))
		SendResponse(c, berror.ErrBind, nil)
		return
	}
	infos, count, err := service.ListUser(r.Username, r.PageNum, r.PageSize)
	if err != nil {
		zap.L().Error("List users error", zap.Error(err))
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, user.ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
