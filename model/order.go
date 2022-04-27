package model

type Order struct {
	Base
	UserID     uint64  `json:"user_id" gorm:"not null" binding:"required" validate:"gte=2"`
	Books      []*Book `json:"books" gorm:"many2many:book_order"`
	OrderPrice float64 `json:"order_price"`
}
