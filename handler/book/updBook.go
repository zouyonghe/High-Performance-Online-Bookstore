package book

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/log"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpdBook(c *gin.Context) {
	zap.L().Info("update book information function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
	bookId, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, nil, err)
		return
	}

	var b model.BookModel

	if err := c.ShouldBindJSON(&b); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBind, nil)
		return
	}
	b.ID = bookId

	if err := b.Validate(); err != nil {
		log.ErrValidate(err)
		SendResponse(c, berror.ErrValidation, nil)
		return
	}

	if err := b.UpdateBook(); err != nil {
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
