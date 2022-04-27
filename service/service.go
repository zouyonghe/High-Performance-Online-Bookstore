package service

import (
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/token"
	"High-Performance-Online-Bookstore/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

func ListUser(username string, pageNum, pageSize int) ([]*model.UserInfo, int64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, pageNum, pageSize)
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
		go func(u *model.User) {
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

func ListBook(title string, pageNum int, pageSize int) ([]*model.BookInfo, int64, error) {
	infos := make([]*model.BookInfo, 0)
	books, count, err := model.ListBook(title, pageNum, pageSize)
	if err != nil {
		return nil, count, err
	}
	var ids []uint64
	for _, b := range books {
		ids = append(ids, b.ID)
	}

	wg := sync.WaitGroup{}
	bookList := model.BookList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.BookInfo, len(books)),
	}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, b := range books {
		wg.Add(1)
		go func(b *model.Book) {
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}
			bookList.Lock.Lock()
			defer bookList.Lock.Unlock()

			bookList.IdMap[b.ID] = &model.BookInfo{
				Id:          b.ID,
				ShortId:     shortId,
				Title:       b.Title,
				Price:       b.Price,
				PublishDate: b.PublishDate,
				Category:    b.Category,
				Author:      b.Author,
				IsSell:      b.IsSell,
				Number:      b.Number,
				CreatedAt:   b.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   b.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(b)
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
		infos = append(infos, bookList.IdMap[id])
	}

	return infos, count, nil
}

func GetIDByParam(c *gin.Context) (uint64, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.ErrConv(err)
		return 0, err
	}
	return uint64(id), nil
}

func GetIDByToken(c *gin.Context) (uint64, error) {
	ctx, err := token.ParseRequest(c)
	if err != nil {
		return 0, err
	}
	return ctx.ID, nil
}
