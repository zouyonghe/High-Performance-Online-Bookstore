package user

import (
	"Jinshuzhai-Bookstore/model"
)

// swagger struct

type SwaggerListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

//--- common user struct ---

// CreateRequest includes username and password
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateResponse includes userId and username
type CreateResponse struct {
	UserId   uint64 `json:"userId"`
	Username string `json:"username"`
}

// LoginRequest includes username and password
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId uint64 `json:"userId"`
	Token  string `json:"token"`
}

type SelfUpdRequest UpdateRequest

type SelfUpdResponse UpdateResponse

type SelfDelRequest struct{}

type SelfDelResponse struct {
	UserId uint64 `json:"userId"`
}

//--- admin user struct ---

// UpdateRequest include username and password to set
type UpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UpdateResponse include userId and username
type UpdateResponse struct {
	UserId   uint64 `json:"userId"`
	Username string `json:"username"`
}

type DeleteRequest struct{}

type DeleteResponse struct {
	UserId uint64 `json:"userId"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount int64             `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
