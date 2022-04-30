package service

import (
	. "High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/constvar"
	"High-Performance-Online-Bookstore/util"
	"sync"
)

func ListBookInfo(title string, pageNum int, pageSize int) ([]*model.BookInfo, error) {
	infos := make([]*model.BookInfo, 0)
	books, err := ListBook(title, pageNum, pageSize)

	if err != nil {
		return nil, err
	}
	var ids []uint64
	for _, b := range books {
		ids = append(ids, b.ID)
	}

	wg := sync.WaitGroup{}
	bookList := model.BookList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.BookInfo, len(books)),
	}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, b := range books {
		wg.Add(1)
		go func(b *model.Book) {
			defer wg.Done()

			shortId, err := util.GenShortId()
			if err != nil {
				errChan <- err
				return
			}
			bookList.Lock.Lock()
			defer bookList.Lock.Unlock()

			bookList.IdMap[b.ID] = &model.BookInfo{
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
			}
		}(b)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	for _, id := range ids {
		infos = append(infos, bookList.IdMap[id])
	}

	return infos, nil
}

// ListBook lists books.
func ListBook(title string, pageNum int, pageSize int) ([]*model.Book, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	var bookList []*model.Book
	var err error
	// Check page number format.
	if pageNum <= 0 {
		pageNum = 1
	}

	offset := (pageNum - 1) * pageSize
	if len(title) > 0 {
		err = DB.Self.Where("title like ?", "%"+title+"%").Offset(offset).Limit(pageSize).Find(&bookList).Error
	} else {
		err = DB.Self.Offset(offset).Limit(pageSize).Find(&bookList).Error
	}

	return bookList, err
}

// ListBookByCategory lists all the books,
// returns book model list,
// count of books and error.
func ListBookByCategory(Category string, pageNum, pageSize int) ([]*model.Book, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	var books []*model.Book
	var count int64
	if err := DB.Self.Where("Category = ?", Category).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&books).Error; err != nil {
		return books, err
	}
	if err := DB.Self.Model(&model.Book{}).Where("Category = ?", Category).Count(&count).Error; err != nil {
		return books, err
	}
	return books, nil
}

// ListBookBySell lists the books on sale.
func ListBookBySell(isSell bool, pageNum, pageSize int) ([]*model.Book, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	var books []*model.Book
	var count int64
	if err := DB.Self.Where("sell = ?", isSell).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&books).Error; err != nil {
		return books, err
	}
	if err := DB.Self.Model(&model.Book{}).Where("on_sale = ?", true).Count(&count).Error; err != nil {
		return books, err
	}
	return books, nil
}

// ListBookBySellAndCategory lists the books
// on sale and specified category.
func ListBookBySellAndCategory(category string, pageNum, pageSize int) ([]*model.Book, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	var books []*model.Book
	var count int64
	if err := DB.Self.Where("category = ?", category).Where("sell = ?", true).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&books).Error; err != nil {
		return books, err
	}
	if err := DB.Self.Model(&model.Book{}).Where("category = ?", category).Where("on_sale = ?", true).Count(&count).Error; err != nil {
		return books, err
	}
	return books, nil
}
