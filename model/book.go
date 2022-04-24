package model

import (
	. "Jinshuzhai-Bookstore/database"
	"Jinshuzhai-Bookstore/pkg/constvar"
	"errors"
	"gorm.io/gorm"
)

// BookBaseModel represents book base information.
type BookBaseModel struct {
	Author      string  `json:"author" gorm:"column:author;not null" binding:"required" validate:"min=1,max=32"`
	Price       float64 `json:"price" gorm:"column:price;not null" binding:"required" validate:"gte=0"`
	PublishDate string  `json:"publish_date" gorm:"column:publishDate;not null" binding:"required" validate:"min=1,max=32,datetime=2006-01-02"`
	Category    string  `json:"category" gorm:"column:category;not null" binding:"required" validate:"min=1,max=32"`
}

// BookModel represents a book information model.
type BookModel struct {
	BaseModel
	//BookBaseModel
	Title  string    `json:"title" gorm:"column:title;not null" binding:"required" validate:"min=1,max=32"`
	Shop   ShopModel `json:"shop" gorm:"column:shop;not null;foreignkey:ShopID" binding:"required"`
	IsSell bool      `json:"sell" gorm:"column:isSale;not null;default:false" binding:"required"`
	Number uint64    `json:"number" gorm:"column:number;not null;default:0" binding:"required" validate:"gte=0"`
}

// TableName returns the table name.
func (b *BookModel) TableName() string {
	return "tb_books"
}

// CreateBook creates a book information.
func (b *BookModel) CreateBook(deleted bool) error {
	if deleted == true {
		bm := &BookModel{}
		DB.Self.Unscoped().Where("title = ?", b.Title).First(&bm)
		DB.Self.Unscoped().Delete(&bm)
	}
	return DB.Self.Create(&b).Error
}

// DeleteBook deletes book information by the book ID.
func DeleteBook(id uint) error {
	return DB.Self.Where("id = ?", id).Delete(&BookModel{}).Error
}

// UpdateBook updates book information.
func (b *BookModel) UpdateBook() error {
	return DB.Self.Save(&b).Error
}

// GetBook gets a book by the book name
// returns book model, deleted and error
func GetBook(title string) (bm *BookModel, deleted bool, err error) {
	bm = &BookModel{}
	d1 := DB.Self.Where("title = ?", title).First(&bm)

	// found record
	if err := d1.Error; err == nil {
		return bm, false, err
	}
	d2 := DB.Self.Unscoped().Where("title = ?", title).First(&bm)
	if errors.Is(d2.Error, gorm.ErrRecordNotFound) {
		return bm, false, gorm.ErrRecordNotFound
	}
	// found record but deleted
	if errors.Is(d1.Error, gorm.ErrRecordNotFound) && d2.Error == nil {
		return bm, true, nil
	}
	return bm, false, nil
}

// GetBookByID gets a book by the book ID.
func GetBookByID(id uint) (*BookModel, error) {
	bm := &BookModel{}
	return bm, DB.Self.Where("id = ?", id).First(&bm).Error
}

// ListBooks lists all books.
func ListBooks(pageNum, pageSize int) ([]*BookModel, int64, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	var books []*BookModel
	var count int64
	if pageSize > 0 && pageNum > 0 {
		offset := (pageNum - 1) * pageSize
		DB.Self.Offset(offset).Limit(pageSize).Find(&books)
		DB.Self.Model(&BookModel{}).Count(&count)
	}
	return books, count, nil
}

// ListBookByCategory lists all the books,
// returns book model list,
// count of books and error.
func ListBookByCategory(Category string, pageNum, pageSize int) ([]*BookModel, int64, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	var books []*BookModel
	var count int64
	if err := DB.Self.Where("Category = ?", Category).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&books).Error; err != nil {
		return books, count, err
	}
	if err := DB.Self.Model(&BookModel{}).Where("Category = ?", Category).Count(&count).Error; err != nil {
		return books, count, err
	}
	return books, count, nil
}

// ListBookBySell lists the books on sale.
func ListBookBySell(pageNum, pageSize int) ([]*BookModel, int64, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	var books []*BookModel
	var count int64
	if err := DB.Self.Where("sell = ?", true).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&books).Error; err != nil {
		return books, count, err
	}
	if err := DB.Self.Model(&BookModel{}).Where("on_sale = ?", true).Count(&count).Error; err != nil {
		return books, count, err
	}
	return books, count, nil
}

// ListBookBySellAndCategory lists the books
// on sale and specified category.
func ListBookBySellAndCategory(category string, pageNum, pageSize int) ([]*BookModel, int64, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	var books []*BookModel
	var count int64
	if err := DB.Self.Where("category = ?", category).Where("sell = ?", true).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&books).Error; err != nil {
		return books, count, err
	}
	if err := DB.Self.Model(&BookModel{}).Where("category = ?", category).Where("on_sale = ?", true).Count(&count).Error; err != nil {
		return books, count, err
	}
	return books, count, nil
}

// SetBookName sets the book name
// and returns error.
func (b *BookModel) SetBookName(name string) error {
	return DB.Self.Model(&b).Update("title", name).Error
}

// SetBookPrice sets the book price
// and returns error.
func (b *BookModel) SetBookPrice(price float64) error {
	return DB.Self.Model(&b).Update("price", price).Error
}

// SetBookCategory sets the book category
// and returns error.
func (b *BookModel) SetBookCategory(category string) error {
	return DB.Self.Model(&b).Update("category", category).Error
}

// SetBookAuthor sets the book author
// and returns error.
func (b *BookModel) SetBookAuthor(author string) error {
	return DB.Self.Model(&b).Update("author", author).Error
}

//SetBookNum sets the book number
//and returns error.
func (b *BookModel) SetBookNum(num int) error {
	return DB.Self.Model(&b).Update("num", num).Error
}

// SetSell sets the book on sale.
func (b *BookModel) SetSell() error {
	return DB.Self.Model(&b).Update("isSale", true).Error
}

// SetUnSell sets the book off sale.
func (b *BookModel) SetUnSell() error {
	return DB.Self.Model(&b).Update("isSale", false).Error
}

// SetBookSell sets the book sale status
// and returns error.
func (b *BookModel) SetBookSell(sell bool) error {
	return DB.Self.Model(&b).Update("isSell", sell).Error
}
