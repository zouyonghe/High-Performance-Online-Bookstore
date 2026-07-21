package service

import (
	. "High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/constvar"
	"High-Performance-Online-Bookstore/util"
)

// ListUserInfo lists user information matching the username.
// The password hash is never included in the result.
func ListUserInfo(username string, pageNum, pageSize int) ([]*model.UserInfo, error) {
	users, err := ListUser(username, pageNum, pageSize)
	if err != nil {
		return nil, err
	}

	infos := make([]*model.UserInfo, 0, len(users))
	for _, u := range users {
		shortId, err := util.GenShortId()
		if err != nil {
			return nil, err
		}
		infos = append(infos, &model.UserInfo{
			Id:        u.ID,
			Username:  u.Username,
			ShortId:   shortId,
			Role:      u.Role,
			CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return infos, nil
}

// ListUser lists all users.
func ListUser(username string, pageNum, pageSize int) ([]*model.User, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize

	userList := make([]*model.User, 0)
	query := DB.Self.Offset(offset).Limit(pageSize)
	if len(username) > 0 {
		query = query.Where("username like ?", "%"+username+"%")
	}
	if err := query.Find(&userList).Error; err != nil {
		return userList, err
	}
	return userList, nil
}
