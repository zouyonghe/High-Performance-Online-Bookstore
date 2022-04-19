package model

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type BaseModel struct {
	ID        uint64         `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time      `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updatedAt" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	ShortId   string `json:"ShortId"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
