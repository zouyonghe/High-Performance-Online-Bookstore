package cart

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	log.AddCartCalled(c)

	// get user id
	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	// bind request body
	var r AddCartRequest
	if err = c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
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

	// set book in cart model
	book, err := model.GetBookByID(r.BookID)
	if err != nil {
		SendError(c, err)
		return
	}
	cb := model.CartBook{
		CartID:    cart.ID,
		BookID:    r.BookID,
		UnitPrice: book.Price,
		Number:    r.Number,
	}

	// add book to cart
	if err = cart.AddBook(cb); err != nil {
		log.ErrAddCart(err)
		SendError(c, err)
		return
	}
	// get book number
	number, err := model.GetBookNumberInCart(cart.ID, book.ID)
	if err != nil {
		log.ErrGetBookNumber(err)
		SendError(c, err)
		return
	}

	rsp := AddCartResponse{
		BookID: r.BookID,
		Number: number,
	}
	SendResponse(c, nil, rsp)
}
