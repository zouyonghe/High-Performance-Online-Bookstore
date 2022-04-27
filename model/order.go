package model

type BaseOrderModel struct {
	BaseModel
	UserID     uint64    `json:"userId" gorm:"column:userId;not null" binding:"required" validate:"min=1,max=32"`
	BookID     BookModel `json:"bookList" gorm:"not null" binding:"required"`
	Number     uint      `json:"number" gorm:"not null" binding:"required" validate:"gte=1"`
	TotalPrice float64   `json:"totalPrice" gorm:"not null"`
}

type OrderModel struct {
	Orders  []BaseOrderModel `json:"orders" gorm:"not null"`
	OrderID uint64           `json:"orderId" gorm:"column:orderId;primary_key"`
}
