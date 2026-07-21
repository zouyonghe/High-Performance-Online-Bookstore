package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// Delete deletes a book.
//
// @Summary Delete a book
// @Description Delete a book specified by book ID (seller or admin only)
// @Tags book
// @Produce json
// @Param id path uint64 true "the ID of the book"
// @Success 200 {object} book.SwaggerDeleteResponse
// @Router /book/{id} [delete]
// @Security ApiKeyAuth
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
