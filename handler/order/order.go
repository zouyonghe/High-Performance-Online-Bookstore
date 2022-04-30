package order

import "High-Performance-Online-Bookstore/model"

type ShowOrderRequest struct{}

type ShowOrderResponse struct {
	OrderPrice float64           `json:"orderPrice"`
	Books      []*model.BookBase `json:"booksInfo"`
}

type CreateOrderRequest struct{}

type CreateOrderResponse struct {
	OrderID    uint64            `json:"orderId"`
	UserID     uint64            `json:"userId"`
	Books      []*model.BookBase `json:"booksInfo"`
	OrderPrice float64           `json:"orderPrice"`
}

type DealOrderRequest struct {
	OrderID    uint64 `json:"orderId"`
	IsApproved bool   `json:"isApproved"`
}

type DealOrderResponse struct {
	OrderID uint64 `json:"orderId"`
}

type ListRequest struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type ListResponse struct {
	TotalCount int                `json:"totalCount"`
	OrderList  []*model.OrderInfo `json:"orderList"`
}
