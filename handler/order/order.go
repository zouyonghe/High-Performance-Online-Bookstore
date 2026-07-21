package order

import (
	"High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/model"
)

type SwaggerCreateOrderResponse struct {
	handler.BaseResponse
	Data CreateOrderResponse `json:"data"`
}

type SwaggerDealOrderResponse struct {
	handler.BaseResponse
	Data DealOrderResponse `json:"data"`
}

type SwaggerListResponse struct {
	handler.BaseResponse
	Data ListResponse `json:"data"`
}

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

// ListRequest represents an order list query.
// The parameters are passed in the query string.
type ListRequest struct {
	PageNum  int `json:"pageNum"  form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type ListResponse struct {
	TotalCount int                `json:"totalCount"`
	OrderList  []*model.OrderInfo `json:"orderList"`
}
