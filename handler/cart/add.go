package cart

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	log.AddCartCalled(c)

	// get user id
	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}
	var r AddCartRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBind, nil)
		return
	}

	// get cart
	cart, err := model.GetCart(userID)
	if err != nil {
		log.ErrGetCart(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}

	book, err := model.GetBookByID(r.BookID)
	if err != nil {
		SendResponse(c, berror.ErrGetBook, nil)
		return
	}

	cb := model.CartBook{
		CartID:    cart.ID,
		BookID:    r.BookID,
		UnitPrice: book.Price,
		Number:    r.Number,
	}

	if err := cart.AddBook(cb); err != nil {
		log.ErrAddCart(err)
	}
	rcb := model.GetCartBook(cart.ID, r.BookID)
	rsp := AddCartResponse{
		BookID: r.BookID,
		Number: rcb.Number,
	}
	SendResponse(c, nil, rsp)
}
