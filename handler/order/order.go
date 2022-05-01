package order

import "High-Performance-Online-Bookstore/model"

type ShowOrderRequest struct{}

type ShowOrderResponse struct {
	OrderPrice float64           `json:"orderPrice"`
	Books      []*model.BookBase `json:"booksInfo"`
}

type CreateOrderRequest struct{}

type CreateOrderResponse struct {
	OrderID uint64 `json:"orderId"`
}

type DealOrderRequest struct {
	OrderID   uint64 `json:"orderId"`
	Operation string `json:"operation"`
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
