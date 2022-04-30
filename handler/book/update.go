package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	log.UpdateBookCalled(c)

	BookID, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, berror.ErrParseToken, nil)
		return
	}

	var b model.Book
	if err = c.ShouldBindJSON(&b); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBindRequest, nil)
		return
	}
	b.ID = BookID

	if err = b.Validate(); err != nil {
		log.ErrValidate(err)
		SendResponse(c, berror.ErrValidation, nil)
		return
	}

	if err = b.UpdateBook(); err != nil {
		log.ErrUpdateBook(err)
		SendResponse(c, berror.ErrDatabase, nil)
		return
	}

	rsp := UpdateResponse{
		BookID:  b.ID,
		Message: "book <" + b.Title + "> update",
	}
	SendResponse(c, nil, rsp)
	return
}
