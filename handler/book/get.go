package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// Get gets a book by ID.
//
// @Summary Get a book
// @Description Get a book specified by book ID
// @Tags book
// @Produce json
// @Param id path uint64 true "the ID of the book"
// @Success 200 {object} book.SwaggerGetResponse
// @Router /book/{id} [get]
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
