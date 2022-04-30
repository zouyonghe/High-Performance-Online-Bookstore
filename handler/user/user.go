package user

import (
	"High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/model"
)

type BaseResponse handler.BaseResponse

// swagger struct

type SwaggerRegisterResponse struct {
	BaseResponse
	Data RegisterResponse `json:"data"`
}

type SwaggerLoginResponse struct {
	BaseResponse
	Data LoginResponse `json:"data"`
}

type SwaggerSelfUpdResponse struct {
	BaseResponse
	Data SelfUpdResponse `json:"data"`
}

type SwaggerSelfDelResponse struct {
	BaseResponse
	Data SelfDelResponse `json:"data"`
}

type SwaggerUpdateResponse struct {
	BaseResponse
	Data UpdateResponse `json:"data"`
}

type SwaggerGetResponse struct {
	BaseResponse
	Data GetResponse `json:"data"`
}

type SwaggerDeleteResponse struct {
	BaseResponse
	Data DeleteResponse `json:"data"`
}

type SwaggerListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

//--- common user struct ---

// RegisterRequest includes username and password
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterResponse includes UserID and username
type RegisterResponse struct {
	UserID   uint64 `json:"UserID"`
	Username string `json:"username"`
}

// LoginRequest includes username and password
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID uint64 `json:"UserID"`
	Token  string `json:"token"`
}

type SelfUpdRequest UpdateRequest

type SelfUpdResponse UpdateResponse

type SelfDelRequest struct{}

type SelfDelResponse struct {
	UserID uint64 `json:"UserID"`
}

//--- admin user struct ---

// UpdateRequest include username and password to set
type UpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UpdateResponse include UserID and username
type UpdateResponse struct {
	UserID   uint64 `json:"UserID"`
	Username string `json:"username"`
}

type GetRequest struct{}

type GetResponse struct {
	UserID   uint64 `json:"UserID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type DeleteRequest struct{}

type DeleteResponse struct {
	UserID  uint64 `json:"UserID"`
	Message string `json:"message"`
}

type ListRequest struct {
	Username string `json:"username"`
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
}

type ListResponse struct {
	TotalCount int64             `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
