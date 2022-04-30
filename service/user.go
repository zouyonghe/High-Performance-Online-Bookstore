package service

import (
	. "High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/constvar"
	"High-Performance-Online-Bookstore/util"
	"sync"
)

func ListUserInfo(username string, pageNum, pageSize int) ([]*model.UserInfo, error) {
	infos := make([]*model.UserInfo, 0)
	users, err := ListUser(username, pageNum, pageSize)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, nil
}

// ListUser lists all users.
func ListUser(username string, pageNum, pageSize int) ([]*model.User, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	offset := (pageNum - 1) * pageSize

	userList := make([]*model.User, 0)
	var count int64
	if len(username) > 0 {
		DB.Self.Where("username like ?", "%"+username+"%").Count(&count)
		if err := DB.Self.Where("username like ?", "%"+username+"%").Offset(offset).Limit(pageSize).Find(&userList).Error; err != nil {
			return userList, err
		}
	} else {
		if err := DB.Self.Offset(offset).Limit(pageSize).Find(&userList).Error; err != nil {
			return userList, err
		}
	}
	return userList, nil
}
