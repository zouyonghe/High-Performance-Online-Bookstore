package admin

import (
	"High-Performance-Online-Bookstore/handler/user"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)
import . "High-Performance-Online-Bookstore/handler"

// List lists users account by specified username format.
//
// @Summary List users account by specified username format.
// @Description List users account by specified username format include id, username, role, etc. The password hash is never returned.
// @Tags user/admin
// @Produce json
// @Param username query string false "fuzzy username"
// @Param pageNum query int false "page number, default 1"
// @Param pageSize query int false "page size, default 10"
// @Success 200 {object} user.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":1,"username":"admin","ShortId":"5P9Ia4QnR","role":"admin","createdAt":"2021-04-18 15:40:33","updatedAt":"2021-04-18 15:40:33"}]}}"
// @Router /user/admin [get]
// @Security ApiKeyAuth
func List(c *gin.Context) {
	log.ListUserCalled(c)

	var r user.ListRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		log.ErrBind(err)
		SendError(c, err)
		return
	}

	infos, err := service.ListUserInfo(r.Username, r.PageNum, r.PageSize)
	if err != nil {
		log.ErrListUsers(err)
		SendError(c, err)
		return
	}

	SendResponse(c, nil, user.ListResponse{
		TotalCount: len(infos),
		UserList:   infos,
	})
}
