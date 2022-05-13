package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	log.GetBookCalled(c)
	BookID, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	book, err := model.GetBookByID(BookID)
	if err != nil {
		log.ErrGetBook(err)
		SendError(c, err)
		return
	}
	SendResponse(c, nil, book)
	return
}
