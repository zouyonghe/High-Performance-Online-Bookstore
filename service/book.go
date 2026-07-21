package service

import (
	. "High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/constvar"
	"High-Performance-Online-Bookstore/util"
)

// ListBookInfo lists book information matching the title.
func ListBookInfo(title string, pageNum int, pageSize int) ([]*model.BookInfo, error) {
	books, err := ListBook(title, pageNum, pageSize)
	if err != nil {
		return nil, err
	}

	infos := make([]*model.BookInfo, 0, len(books))
	for _, b := range books {
		shortId, err := util.GenShortId()
		if err != nil {
			return nil, err
		}
		infos = append(infos, &model.BookInfo{
			Id:          b.ID,
			ShortId:     shortId,
			Title:       b.Title,
			Price:       b.Price,
			PublishDate: b.PublishDate,
			Category:    b.Category,
			Author:      b.Author,
			IsSell:      b.IsSell,
			Number:      b.Number,
			CreatedAt:   b.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   b.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return infos, nil
}

// ListBook lists books.
func ListBook(title string, pageNum int, pageSize int) ([]*model.Book, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize

	bookList := make([]*model.Book, 0)
	query := DB.Self.Offset(offset).Limit(pageSize)
	if len(title) > 0 {
		query = query.Where("title like ?", "%"+title+"%")
	}
	if err := query.Find(&bookList).Error; err != nil {
		return bookList, err
	}
	return bookList, nil
}

// ListBookByCategory lists the books of the specified category.
func ListBookByCategory(category string, pageNum, pageSize int) ([]*model.Book, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	books := make([]*model.Book, 0)
	err := DB.Self.Where("category = ?", category).
		Offset((pageNum - 1) * pageSize).Limit(pageSize).
		Find(&books).Error
	return books, err
}

// ListBookBySell lists the books on sale.
func ListBookBySell(isSell bool, pageNum, pageSize int) ([]*model.Book, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	books := make([]*model.Book, 0)
	err := DB.Self.Where("sell = ?", isSell).
		Offset((pageNum - 1) * pageSize).Limit(pageSize).
		Find(&books).Error
	return books, err
}

// ListBookBySellAndCategory lists the books
// on sale and specified category.
func ListBookBySellAndCategory(category string, pageNum, pageSize int) ([]*model.Book, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	books := make([]*model.Book, 0)
	err := DB.Self.Where("category = ?", category).Where("sell = ?", true).
		Offset((pageNum - 1) * pageSize).Limit(pageSize).
		Find(&books).Error
	return books, err
}
