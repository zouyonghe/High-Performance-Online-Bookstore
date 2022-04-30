package cart

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	log.ShowCartCalled(c)

	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, berror.ErrParseToken, nil)
		return
	}
	ct, err := model.GetCart(userID)
	if err != nil {
		log.ErrGetCart(err)
		SendResponse(c, berror.ErrGetCart, nil)
		return
	}
	bookList, cartPrice, err := ct.GetBookList()
	if err != nil {
		log.ErrGetCartBookList(err)
		SendResponse(c, berror.ErrGetBookList, nil)
	}
	rsp := ShowCartResponse{
		CartPrice: cartPrice,
		Books:     bookList,
	}
	SendResponse(c, nil, rsp)
}
