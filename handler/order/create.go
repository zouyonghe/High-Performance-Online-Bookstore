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
		SendResponse(c, berror.InternalServerError, nil)
		return
	}
	// get cart
	ct, err := model.GetCart(userID)
	if err != nil {
		log.ErrGetCart(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}

	bookList, cartPrice, err := ct.GetBookList()
	if cartPrice == 0 {
		SendResponse(c, berror.ErrNothingInCart, nil)
		return
	}
	if err != nil {
		log.ErrGetCartBookList(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}
	err = model.CreateOrder(userID)
	if err != nil {
		log.ErrCreateOrder(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}
	o, err := model.GetOrder(userID)
	if err != nil {
		log.ErrGetOrder(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}
	if err = o.AddBook(bookList); err != nil {
		log.ErrAddOrder(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}
	if err = o.UpdateOrderPrice(); err != nil {
		log.ErrUpdateOrderPrice(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}
	SendResponse(c, nil, nil)
}
