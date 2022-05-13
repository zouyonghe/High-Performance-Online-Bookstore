package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	log.ListBookCalled(c)

	var r ListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendError(c, err)
		return
	}
	infos, err := service.ListBookInfo(r.Title, r.PageNum, r.PageSize)
	if err != nil {
		log.ErrListBook(err)
		SendError(c, err)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: len(infos),
		BookList:   infos,
	})
}
