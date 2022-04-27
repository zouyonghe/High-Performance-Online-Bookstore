package model

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}

// BaseModel represents a base model.
type BaseModel struct {
	ID        uint64         `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time      `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updatedAt" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt" sql:"index" json:"-"`
}

// User models

// UserInfo represents user information.
type UserInfo struct {
	Id        uint64 `json:"userId"`
	Username  string `json:"username"`
	ShortId   string `json:"ShortId"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// UserList represents users list.
type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// Book models

// BookInfo represents book information.
type BookInfo struct {
	Id          uint64  `json:"bookId"`
	Title       string  `json:"title"`
	ShortId     string  `json:"shortId"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	PublishDate string  `json:"publishDate"`
	Category    string  `json:"category"`
	IsSell      bool    `json:"isSell"`
	Number      uint64  `json:"number"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

// BookClass represents the class of a book.
type BookClass struct {
	ClassId   uint64 `json:"classId"`
	ClassName string `json:"className"`
}

// BookList represents books list.
type BookList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*BookInfo
}

type CartInfo struct {
}

type OrderInfo struct {
}
