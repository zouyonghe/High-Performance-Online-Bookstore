package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	log.DelBookCalled(c)
	BookID, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	bm, err := model.GetBookByID(BookID)
	if err != nil {
		log.ErrGetBook(err)
		SendError(c, err)
		return
	}

	title := bm.Title

	if err = model.DeleteBook(BookID); err != nil {
		log.ErrDelBook(err)
		SendError(c, err)
		return
	}

	rsp := DeleteResponse{
		BookID:  BookID,
		Message: "Book <" + title + "> delete",
	}
	SendResponse(c, nil, rsp)
}
