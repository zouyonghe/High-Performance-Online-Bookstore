package cart

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	log.ShowCartCalled(c)

	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}
	ct, err := model.GetCart(userID)
	if err != nil {
		log.ErrGetCart(err)
		SendError(c, err)
		return
	}
	bookList, cartPrice, err := ct.GetBookList()
	if err != nil {
		log.ErrGetCartBookList(err)
		SendError(c, err)
		return
	}
	rsp := ShowCartResponse{
		CartPrice: cartPrice,
		Books:     bookList,
	}
	SendResponse(c, nil, rsp)
}
