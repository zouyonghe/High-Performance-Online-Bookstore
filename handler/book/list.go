package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// List lists books by title and category.
//
// @Summary List books
// @Description List books by fuzzy title and exact category with pagination
// @Tags book
// @Produce json
// @Param title query string false "fuzzy book title"
// @Param category query string false "exact book category"
// @Param pageNum query int false "page number, default 1"
// @Param pageSize query int false "page size, default 10"
// @Success 200 {object} book.SwaggerListResponse
// @Router /book [get]
func List(c *gin.Context) {
	log.ListBookCalled(c)

	var r ListRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		log.ErrBind(err)
		SendError(c, err)
		return
	}
	infos, err := service.ListBookInfo(r.Title, r.Category, r.PageNum, r.PageSize)
	if err != nil {
		log.ErrListBook(err)
		SendError(c, err)
		return
	}
	rsp := ListResponse{
		TotalCount: len(infos),
		BookList:   infos,
	}
	SendResponse(c, nil, rsp)
}
