package model

import (
	. "High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/log"
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
	OrderID   uint64  `json:"CartID" gorm:"primaryKey"`
	BookID    uint64  `json:"BookID" gorm:"primaryKey"`
	UnitPrice float64 `json:"unitPrice"`
	Number    uint    `json:"Number"`
}

// TableName returns the table name.
func (o *Order) TableName() string {
	return "tb_orders"
}

func CreateOrder(userID uint64) error {
	return DB.Self.Create(&Order{
		UserID:     userID,
		Books:      make([]OrderBook, 0),
		OrderPrice: 0,
		IsApproved: false,
	}).Error
}

func GetOrder(userID uint64) (*Order, error) {
	var order Order
	err := DB.Self.Where("user_id = ?", userID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *Order) AddBook(books []*BookBase) error {
	for _, book := range books {
		b, deleted, err := GetBook(book.Title)
		if err != nil || deleted == true {
			log.ErrGetBook(err)
		}
		o.Books = append(o.Books, OrderBook{
			OrderID:   o.ID,
			BookID:    b.ID,
			UnitPrice: b.Price,
			Number:    book.Number,
		})
	}
	return DB.Self.Save(&o).Error
}

func (o *Order) TotalPrice() float64 {
	var totalPrice float64
	var ob []*OrderBook
	if err := DB.Self.Model(&OrderBook{}).Where("cart_id = ?", o.ID).Find(&ob).Error; err != nil {
		log.ErrGetOrderBook(err)
	}
	for _, book := range ob {
		totalPrice += book.UnitPrice * float64(book.Number)
	}
	return totalPrice
}

func (o *Order) UpdateOrderPrice() error {
	o.OrderPrice = o.TotalPrice()
	return DB.Self.Save(&o).Error
}
