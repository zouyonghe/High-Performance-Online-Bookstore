package book

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/log"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddBook(c *gin.Context) {
	log.AddBookCalled(c)

	var r AddRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		zap.L().Error("AddBook Bind", zap.Error(err))
		SendResponse(c, berror.ErrBind, nil)
		return
	}

	b := model.BookModel{
		Title:       r.Title,
		Author:      r.Author,
		Price:       r.Price,
		PublishDate: r.PublishDate,
		Category:    r.Category,
		IsSell:      r.IsSell,
		Number:      r.Number,
	}

	// Validate the data.
	if err := b.Validate(); err != nil {
		log.ErrValidate(err)
		SendResponse(c, err, nil)
		return
	}

	// Validate if the book is exists.
	_, deleted, err := model.GetBook(r.Title)
	// If the book exists and deleted is false, send an error.
	if deleted == false && err == nil {
		log.ErrBookExists()
		SendResponse(c, berror.ErrBookExists, nil)
		return
	}

	// Insert the book into the database.
	if err := b.CreateBook(deleted); err != nil {
		log.ErrCreateBook(err)
		SendResponse(c, err, nil)
		return
	}

	rsp := AddResponse{
		BookID: b.ID,
		Title:  b.Title,
	}

	// Return the book ID and title of the book.
	SendResponse(c, nil, rsp)
}
