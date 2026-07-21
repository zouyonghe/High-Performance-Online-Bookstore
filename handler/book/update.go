package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// Update updates a book.
//
// @Summary Update a book
// @Description Update a book specified by book ID (seller or admin only)
// @Tags book
// @Accept json
// @Produce json
// @Param id path uint64 true "the ID of the book"
// @Param book body book.AddRequest true "book information"
// @Success 200 {object} book.SwaggerUpdateResponse
// @Router /book/{id} [put]
// @Security ApiKeyAuth
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
