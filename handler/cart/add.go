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

	// bind request body
	var r AddCartRequest
	if err = c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBindRequest, nil)
		return
	}

	// get cart
	cart, err := model.GetCart(userID)
	if err != nil {
		log.ErrGetCart(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}

	// set book in cart model
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

	// add book to cart
	if err = cart.AddBook(cb); err != nil {
		log.ErrAddCart(err)
		SendResponse(c, err, nil)
		return
	}
	// get book number
	number, err := model.GetBookNumberInCart(cart.ID, book.ID)
	if err != nil {
		log.ErrGetBookNumber(err)
		SendResponse(c, err, nil)
	}

	rsp := AddCartResponse{
		BookID: r.BookID,
		Number: number,
	}
	SendResponse(c, nil, rsp)
}
