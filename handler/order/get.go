package order

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// Get gets one order.
//
// @Summary Get one order
// @Description Get one order by ID for manager roles
// @Tags order
// @Produce json
// @Param id path uint64 true "the ID of the order"
// @Success 200 {object} order.SwaggerGetResponse
// @Router /order/{id} [get]
// @Security ApiKeyAuth
func Get(c *gin.Context) {
	orderID, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	o, err := model.GetOrder(orderID)
	if err != nil {
		log.ErrGetOrder(err)
		SendError(c, err)
		return
	}

	books, err := o.GetOrderBooks()
	if err != nil {
		log.ErrGetOrderBook(err)
		SendError(c, err)
		return
	}

	bookList := make([]model.OrderBook, 0, len(books))
	for _, b := range books {
		bookList = append(bookList, *b)
	}

	SendResponse(c, nil, GetResponse{Order: model.OrderInfo{
		OrderID:    o.ID,
		UserID:     o.UserID,
		Books:      bookList,
		OrderPrice: o.OrderPrice,
		CreatedAt:  o.CreatedAt.Format("2006-01-02 15:04:05"),
		Status:     o.Status,
	}})
}
