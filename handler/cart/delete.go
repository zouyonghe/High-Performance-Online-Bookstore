package cart

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	log.DeleteCartCalled(c)

	// bind request body
	var r DeleteCartRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendError(c, err)
		return
	}

	// get user id
	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	// get cart
	cart, err := model.GetCart(userID)
	if err != nil {
		log.ErrGetCart(err)
		SendError(c, err)
		return
	}

	// delete books from cart
	if err = model.DeleteFromCart(cart.ID, r.BookID, r.Number); err != nil {
		log.ErrDeleteCart(err)
		SendError(c, err)
		return
	}

	rcb := model.GetCartBook(cart.ID, r.BookID)
	rsp := AddCartResponse{
		BookID: r.BookID,
		Number: rcb.Number,
	}
	SendResponse(c, nil, rsp)
}
