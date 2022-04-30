package model

import (
	. "High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/pkg/berror"
)

type Order struct {
	Base
	UserID     uint64      `json:"userID" gorm:"not null" binding:"required" validate:"gte=2"`
	Books      []OrderBook `json:"books"`
	OrderPrice float64     `json:"orderPrice"`
	IsApproved bool        `json:"isApproved" gorm:"not null;default:false"`
}

type OrderBook struct {
	Base
	OrderID   uint64  `json:"-" gorm:"primaryKey"`
	BookID    uint64  `json:"BookID" gorm:"primaryKey"`
	UnitPrice float64 `json:"unitPrice"`
	Number    uint    `json:"Number"`
}

// TableName returns the table name.
func (o *Order) TableName() string {
	return "tb_orders"
}

func CreateOrder(userID uint64) (*Order, error) {
	o := &Order{
		UserID:     userID,
		Books:      make([]OrderBook, 0),
		OrderPrice: 0,
		IsApproved: false,
	}
	return o, DB.Self.Create(o).Error
}

func GetOrder(orderID uint64) (*Order, error) {
	var order Order
	err := DB.Self.Model(&Order{}).Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *Order) AddBook(books []*CartBook) error {
	for _, book := range books {
		o.Books = append(o.Books, OrderBook{
			OrderID:   o.ID,
			BookID:    book.BookID,
			UnitPrice: book.UnitPrice,
			Number:    book.Number,
		})
	}
	return DB.Self.Save(&o).Error
}

func (o *Order) TotalPrice() float64 {
	var totalPrice float64
	var ob []*OrderBook
	if err := DB.Self.Model(&OrderBook{}).Where("order_id = ?", o.ID).Find(&ob).Error; err != nil {
		log.ErrGetOrderBook(err)
	}
	for _, book := range ob {
		totalPrice += book.UnitPrice * float64(book.Number)
	}
	return totalPrice
}

func (o *Order) SetOrderPrice() error {
	o.OrderPrice = o.TotalPrice()
	return DB.Self.Save(&o).Error
}

func (o *Order) ApproveOrder() error {
	r := DB.Self.Model(&Order{}).Where("id = ?", o.ID).First(&Order{})
	if r.Error != nil {
		return r.Error
	}
	o.IsApproved = true
	var ob []*OrderBook
	if err := DB.Self.Model(&OrderBook{}).Where("order_id = ?", o.ID).Find(&ob).Error; err != nil {
		return err
	}
	for _, book := range ob {
		var b Book
		if err := DB.Self.Model(&Book{}).Where("id = ?", book.BookID).First(&b).Error; err != nil {
			return err
		}
		if b.IsSell == false {
			return berror.ErrBookNotSell
		}
		b.Number -= book.Number
		if b.Number < 0 {
			b.Number = 0
			return berror.ErrBookNotEnough
		} else if b.Number == 0 {
			b.IsSell = false
		}

		if err := b.UpdateBook(); err != nil {
			return err
		}
	}
	return DB.Self.Save(&o).Error
}

func (o *Order) DeleteOrder() error {
	return DB.Self.Delete(&o).Error
}

func (o *Order) GetOrderBooks() ([]*OrderBook, error) {
	var ob []*OrderBook
	if err := DB.Self.Model(&OrderBook{}).Where("order_id = ?", o.ID).Find(&ob).Error; err != nil {
		return nil, err
	}
	return ob, nil
}

func (o *Order) GetOrderBook(bookID uint64) (*OrderBook, error) {
	var ob OrderBook
	if err := DB.Self.Model(&OrderBook{}).Where("order_id = ? and book_id = ?", o.ID, bookID).First(&ob).Error; err != nil {
		return nil, err
	}
	return &ob, nil
}

func (o *Order) UpdateOrderBook(bookID uint64, number uint) error {
	var ob OrderBook
	if err := DB.Self.Model(&OrderBook{}).Where("order_id = ? and book_id = ?", o.ID, bookID).First(&ob).Error; err != nil {
		return err
	}
	ob.Number = number
	return DB.Self.Save(&ob).Error
}

func (o *Order) DeleteOrderBook(bookID uint64) error {
	var ob OrderBook
	if err := DB.Self.Model(&OrderBook{}).Where("order_id = ? and book_id = ?", o.ID, bookID).First(&ob).Error; err != nil {
		return err
	}
	return DB.Self.Delete(&ob).Error
}

func (o *Order) Deal(IsApproved bool) error {
	o.IsApproved = IsApproved
	if IsApproved == true {
		if err := o.ApproveOrder(); err != nil {
			return err
		}
	} else {
		if err := o.DeleteOrder(); err != nil {
			return err
		}
	}
	return nil
}
