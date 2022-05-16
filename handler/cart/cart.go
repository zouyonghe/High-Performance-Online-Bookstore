package cart

import "High-Performance-Online-Bookstore/model"

type ShowCartRequest struct{}

type ShowCartResponse struct {
	CartPrice float64           `json:"cartPrice"`
	Books     []*model.BookBase `json:"booksInfo"`
}

type AddCartRequest struct {
	BookID uint64 `json:"BookID"`
	Number uint   `json:"number"`
}

type AddCartResponse struct {
	BookID uint64 `json:"BookID"`
	Number uint   `json:"Number"`
}

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
