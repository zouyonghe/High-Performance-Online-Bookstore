package model

import (
	. "High-Performance-Online-Bookstore/database"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BookBase struct {
	Index  int     `json:"index"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Number uint    `json:"number"`
}

// Book represents a book information model.
type Book struct {
	Base
	Title       string  `json:"title" gorm:"column:title;not null"             binding:"required"  validate:"min=1,max=32"`
	Price       float64 `json:"price" gorm:"column:price;not null"             binding:"required"  validate:"gte=0"`
	IsSell      bool    `json:"isSell" gorm:"column:is_sell;not null;default:false"`
	Number      uint    `json:"number" gorm:"column:number;not null;default:0" binding:"required"  validate:"gte=0"`
	Author      string  `json:"author" gorm:"column:author;not null"           binding:"required"  validate:"min=5,max=32"`
	PublishDate string  `json:"publishDate" gorm:"column:publish_date;not null" binding:"required"  validate:"datetime=2006-01-02"`
	Category    string  `json:"category" gorm:"column:category;not null"       binding:"required"  validate:"min=1,max=32"`
}

// TableName returns the table name.
func (b *Book) TableName() string {
	return "tb_books"
}

// CreateBook creates a book record.
func (b *Book) CreateBook(deleted bool) error {
	if deleted == true {
		bm := &Book{}
		DB.Self.Unscoped().Where("title = ?", b.Title).First(&bm)
		DB.Self.Unscoped().Delete(&bm)
	}
	return DB.Self.Create(&b).Error
}

// DeleteBook deletes book record by the book ID.
func DeleteBook(id uint64) error {
	return DB.Self.Where("id = ?", id).Delete(&Book{}).Error
}

// UpdateBook updates book record.
func (b *Book) UpdateBook() error {
	return DB.Self.Save(&b).Error
}

// GetBook gets a book by the book name
// returns book model, deleted and error.
func GetBook(title string) (bm *Book, deleted bool, err error) {
	bm = &Book{}
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
func GetBookByID(id uint64) (*Book, error) {
	bm := &Book{}
	return bm, DB.Self.Where("id = ?", id).First(&bm).Error
}

// SetBookName sets the book name
// and returns error.
func (b *Book) SetBookName(name string) error {
	return DB.Self.Model(&b).Update("title", name).Error
}

// SetBookPrice sets the book price
// and returns error.
func (b *Book) SetBookPrice(price float64) error {
	return DB.Self.Model(&b).Update("price", price).Error
}

// SetBookCategory sets the book category
// and returns error.
func (b *Book) SetBookCategory(category string) error {
	return DB.Self.Model(&b).Update("category", category).Error
}

// SetBookAuthor sets the book author
// and returns error.
func (b *Book) SetBookAuthor(author string) error {
	return DB.Self.Model(&b).Update("author", author).Error
}

//SetBookNum sets the book number
//and returns error.
func (b *Book) SetBookNum(num int) error {
	return DB.Self.Model(&b).Update("num", num).Error
}

// SetSell sets the book on sale.
func (b *Book) SetSell() error {
	return DB.Self.Model(&b).Update("isSale", true).Error
}

// SetUnSell sets the book off sale.
func (b *Book) SetUnSell() error {
	return DB.Self.Model(&b).Update("isSale", false).Error
}

// SetBookSell sets the book sale status
// and returns error.
func (b *Book) SetBookSell(sell bool) error {
	return DB.Self.Model(&b).Update("isSell", sell).Error
}

// Validate the fields.
func (b *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
