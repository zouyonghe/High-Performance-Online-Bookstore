package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	log.ListBookCalled(c)

	var r ListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBindRequest, nil)
		return
	}
	infos, count, err := service.ListBookInfo(r.Title, r.PageNum, r.PageSize)
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
