package admin

import (
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sync"
)
import . "Jinshuzhai-Bookstore/handler"

// List lists all users account.
//
// @Summary List all users account
// @Description List all users account include id, username, encrypted password, etc
// @Tags user/admin
// @Produce  json
// @Success 200 {object} user.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":1,"username":"admin","ShortId":"5P9Ia4QnR","password":"$2a$10$Fv9BWzqsiQ.JuuGdcXdvN.Fx3ml.dVR47W22GoJMWQAlm9wHQIMVe","role":"admin","createdAt":"2021-04-18 15:40:33","updatedAt":"2021-04-18 15:40:33"}]}}"
// @Router /user/admin [get]
// @Security ApiKeyAuth
func List(c *gin.Context) {
	zap.L().Info("List function called", zap.String("X-Request-Id", util.GetReqID(c)))
	var r user.ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, berror.ErrBind, nil)
		return
	}

	infos, count, err := listUser(r.Username, r.PageNum, r.PageSize)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, user.ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}

func listUser(username string, offset, limit int) ([]*model.UserInfo, int64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	var ids []uint64
	for _, u := range users {
		ids = append(ids, u.ID)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.ID] = &model.UserInfo{
				Id:        u.ID,
				Username:  u.Username,
				ShortId:   shortId,
				Password:  u.Password,
				Role:      u.Role,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}
