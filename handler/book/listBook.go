package book

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/log"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func ListBook(c *gin.Context) {
	log.ListBookCalled(c)

	var r ListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBind, nil)
		return
	}
	infos, count, err := service.ListBook(r.Title, r.PageNum, r.PageSize)
	if err != nil {
		log.ErrListBooks(err)
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		BookList:   infos,
	})
}
