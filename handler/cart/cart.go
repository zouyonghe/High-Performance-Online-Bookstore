package cart

import (
	"High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/model"
)

type SwaggerShowCartResponse struct {
	handler.BaseResponse
	Data ShowCartResponse `json:"data"`
}

type SwaggerAddCartResponse struct {
	handler.BaseResponse
	Data AddCartResponse `json:"data"`
}

type SwaggerDeleteCartResponse struct {
	handler.BaseResponse
	Data DeleteCartResponse `json:"data"`
}

type SwaggerClearCartResponse struct {
	handler.BaseResponse
	Data ClearCartResponse `json:"data"`
}

type ShowCartRequest struct{}

type ShowCartResponse struct {
	CartPrice float64           `json:"cartPrice"`
	Books     []*model.BookBase `json:"booksInfo"`
}

// AddCartRequest represents a request for adding a book to the cart.
type AddCartRequest struct {
	BookID uint64 `json:"BookID"`
	Number uint   `json:"number"`
}

type AddCartResponse struct {
	BookID uint64 `json:"BookID"`
	Number uint   `json:"Number"`
}

// DeleteCartRequest represents a request for removing books from the cart.
type DeleteCartRequest struct {
	BookID uint64 `json:"BookID"`
	Number uint   `json:"Number"`
}

type DeleteCartResponse struct {
	BookID uint64 `json:"BookID"`
	Number uint   `json:"Number"`
}

type ClearCartRequest struct{}

type ClearCartResponse struct {
	Message string `json:"Message"`
}
