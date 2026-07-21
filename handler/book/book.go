package book

import (
	"High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/model"
)

type SwaggerAddResponse struct {
	handler.BaseResponse
	Data AddResponse `json:"data"`
}

type SwaggerListResponse struct {
	handler.BaseResponse
	Data ListResponse `json:"data"`
}

type SwaggerGetResponse struct {
	handler.BaseResponse
	Data model.Book `json:"data"`
}

type SwaggerUpdateResponse struct {
	handler.BaseResponse
	Data UpdateResponse `json:"data"`
}

type SwaggerDeleteResponse struct {
	handler.BaseResponse
	Data DeleteResponse `json:"data"`
}

// AddRequest represents a request for adding a book.
type AddRequest struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	PublishDate string  `json:"publishDate"`
	Category    string  `json:"category"`
	IsSell      bool    `json:"isSell"`
	Number      uint    `json:"number"`
}

type AddResponse struct {
	BookID uint64 `json:"BookID"`
	Title  string `json:"title"`
}

// ListRequest represents a book list query.
// The parameters are passed in the query string.
type ListRequest struct {
	Title    string `json:"title"    form:"title"`
	Category string `json:"category" form:"category"`
	PageNum  int    `json:"pageNum"  form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type ListResponse struct {
	TotalCount int               `json:"totalCount"`
	BookList   []*model.BookInfo `json:"bookList"`
}

type DeleteRequest struct{}

type DeleteResponse struct {
	BookID  uint64 `json:"BookID"`
	Message string `json:"message"`
}

type UpdateRequest struct{}

type UpdateResponse struct {
	BookID  uint64 `json:"BookID"`
	Message string `json:"message"`
}

type GetRequest struct{}
