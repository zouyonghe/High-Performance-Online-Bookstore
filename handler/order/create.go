package order

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	log.CreateOrderCalled(c)

	// get user id
	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}
	// get cart
	ct, err := model.GetCart(userID)
	if err != nil {
		log.ErrGetCart(err)
		SendError(c, err)
		return
	}
	// get cart book list
	bookList, err := ct.GetCartBook()
	if len(bookList) == 0 {
		SendResponse(c, berror.ErrNothingInCart, nil)
		return
	}
	if err != nil {
		log.ErrGetCartBookList(err)
		SendError(c, err)
		return
	}
	// create order
	o, err := model.CreateOrder(userID)
	if err != nil {
		log.ErrCreateOrder(err)
		SendError(c, err)
		return
	}
	if err = o.AddBook(bookList); err != nil {
		log.ErrAddOrder(err)
		SendError(c, err)
		return
	}
	if err = o.SetOrderPrice(); err != nil {
		log.ErrUpdateOrderPrice(err)
		SendError(c, err)
		return
	}

	// clear cart
	if err = ct.ClearCart(); err != nil {
		log.ErrDeleteCart(err)
		SendResponse(c, berror.ErrClearCart, nil)
		return
	}
	rsp := &CreateOrderResponse{
		OrderID: o.ID,
	}
	SendResponse(c, nil, rsp)
}
