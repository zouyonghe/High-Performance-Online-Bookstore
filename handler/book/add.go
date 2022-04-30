package book

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Add(c *gin.Context) {
	log.AddBookCalled(c)

	var r AddRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		zap.L().Error("AddBook Bind", zap.Error(err))
		SendResponse(c, berror.ErrBindRequest, nil)
		return
	}
	b := model.Book{
		Title:       r.Title,
		Author:      r.Author,
		PublishDate: r.PublishDate,
		Category:    r.Category,
		Price:       r.Price,
		IsSell:      r.IsSell,
		Number:      r.Number,
	}
	// Validate the data.
	if err := b.Validate(); err != nil {
		log.ErrValidate(err)
		SendResponse(c, berror.ErrValidation, nil)
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
	if err = b.CreateBook(deleted); err != nil {
		log.ErrCreateBook(err)
		SendResponse(c, berror.ErrCreateBook, nil)
		return
	}

	rsp := AddResponse{
		BookID: b.ID,
		Title:  b.Title,
	}

	// Return the book ID and title of the book.
	SendResponse(c, nil, rsp)
}
