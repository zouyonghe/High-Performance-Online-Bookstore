package model

import (
	. "High-Performance-Online-Bookstore/database"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	Base
	UserID    uint64     `json:"userID"`
	Books     []CartBook `json:"books"`
	CartPrice float64    `json:"cartPrice"`
}

type CartBook struct {
	Base
	CartID    uint64  `json:"CartID" gorm:"primaryKey"`
	BookID    uint64  `json:"BookID" gorm:"primaryKey"`
	UnitPrice float64 `json:"unitPrice"`
	Number    uint    `json:"Number"`
}

func CreateCart(userID uint64) error {
	return DB.Self.Create(&Cart{
		UserID:    userID,
		Books:     make([]CartBook, 0),
		CartPrice: 0,
	}).Error
}

// GetCart returns a cart model.
// If not exists, create a new one.
func GetCart(userID uint64) (*Cart, error) {
	c := &Cart{}
	r := DB.Self.Where("user_id = ?", userID).First(c)
	if r.Error != nil {
		return nil, r.Error
	}

	return c, nil
}

func (c *Cart) GetBookList() (books []*BookBase, cartPrice float64, err error) {
	cartPrice = 0
	var cb []*CartBook
	books = make([]*BookBase, 0)
	if err := DB.Self.Model(&CartBook{}).Where("cart_id = ?", c.ID).Find(&cb).Error; err != nil {
		return nil, 0, err
	}
	for i, book := range cb {
		b, err := GetBookByID(book.BookID)
		if err != nil {
			return nil, 0, err
		}
		bookInfo := &BookBase{
			Index:  i + 1,
			Title:  b.Title,
			Price:  book.UnitPrice,
			Number: book.Number,
		}
		books = append(books, bookInfo)
		cartPrice += book.UnitPrice * float64(book.Number)
	}
	return books, cartPrice, nil
}

func (c *Cart) AddBook(cb CartBook) error {
	var result CartBook
	r := DB.Self.Where("cart_id = ? AND book_id = ?", cb.CartID, cb.BookID).First(&result)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		c.Books = append(c.Books, cb)
		return DB.Self.Save(c).Error
	} else {
		result.Number += cb.Number
		return DB.Self.Save(&result).Error
	}
}

// GetCartBook returns a CartBook.
func GetCartBook(cartID uint64, bookID uint64) *CartBook {
	var cartBook CartBook
	DB.Self.Where("cart_id = ? AND book_id = ?", cartID, bookID).First(&cartBook)
	return &cartBook
}

func DeleteCart(cartID uint64) error {
	return DB.Self.Where("id = ?", cartID).Delete(&Cart{}).Error
}

func DeleteFromCart(cartID uint64, bookID uint64, number uint) error {
	var cartBook CartBook
	r := DB.Self.Where("cart_id = ? AND book_id = ?", cartID, bookID).First(&cartBook)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return r.Error
	}
	if cartBook.Number <= number {
		return DB.Self.Where("cart_id = ? AND book_id = ?", cartID, bookID).Delete(&CartBook{}).Error
	} else {
		cartBook.Number -= number
		return DB.Self.Save(&cartBook).Error
	}
}

func CheckCartBook(cartID uint64, bookID uint64) bool {
	var cartBook CartBook
	DB.Self.Where("cart_id = ? AND book_id = ?", cartID, bookID).First(&cartBook)
	if cartBook.CartID == 0 {
		return false
	}
	return true
}

func ClearCart(cartID uint64) error {
	return DB.Self.Where("cart_id = ?", cartID).Delete(&CartBook{}).Error
}

func CreateCartByName(userName string) error {
	var user User
	r := DB.Self.Model(&User{}).Where("userName = ?", userName).First(&user)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return r.Error
	}
	return CreateCart(user.ID)
}
