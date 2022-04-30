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
		SendResponse(c, nil, err)
		return
	}

	bm, err := model.GetBookByID(BookID)
	if err != nil {
		log.ErrGetBook(err)
		c.JSON(400, gin.H{
			"message": "Book not found",
		})
		return
	}

	title := bm.Title

	if err := model.DeleteBook(BookID); err != nil {
		log.ErrDelBook(err)
		SendResponse(c, err, nil)
		return
	}

	rsp := DeleteResponse{
		BookID:  BookID,
		Message: "Book <" + title + "> delete",
	}
	SendResponse(c, nil, rsp)
	return
}
