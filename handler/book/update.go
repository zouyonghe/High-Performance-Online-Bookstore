package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	log.UpdateBookCalled(c)

	BookID, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	var b model.Book
	if err = c.ShouldBindJSON(&b); err != nil {
		log.ErrBind(err)
		SendError(c, err)
		return
	}
	b.ID = BookID

	if err = b.Validate(); err != nil {
		log.ErrValidate(err)
		SendError(c, err)
		return
	}

	if err = b.UpdateBook(); err != nil {
		log.ErrUpdateBook(err)
		SendError(c, err)
		return
	}

	rsp := UpdateResponse{
		BookID:  b.ID,
		Message: "book <" + b.Title + "> update",
	}
	SendResponse(c, nil, rsp)
	return
}
