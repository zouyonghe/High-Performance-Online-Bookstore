package model

type Cart struct {
	Base
	UserId    int64   `json:"user_id"`
	Books     []*Book `json:"books" gorm:"many2many:book_cart"`
	CartPrice float64 `json:"cart_price"`
}
