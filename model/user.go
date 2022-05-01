package model

import (
	. "High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/pkg/auth"
	"High-Performance-Online-Bookstore/pkg/berror"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UserBaseModel represents base user information.
type UserBaseModel struct {
	Phone   string `json:"phone"   gorm:"column:phone;default:"   validate:"min=5,max=32"`
	Address string `json:"address" gorm:"column:address;default:" validate:"min=5,max=128"`
}

// User represents user information.
type User struct {
	Base
	Username string  `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=2,max=32"`
	Password string  `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=32"`
	Role     string  `json:"role"     gorm:"column:role;not null;default:general"        validate:"oneof=general seller admin"`
	Cart     Cart    `json:"-"`
	Orders   []Order `json:"-"`
}

// TableName returns the table name.
func (u *User) TableName() string {
	return "tb_users"
}

// CreateUser creates a new user account.
func (u *User) CreateUser(deleted bool) error {
	if deleted == true {
		um := &User{}
		DB.Self.Unscoped().Where("username = ?", u.Username).First(&um)
		DB.Self.Unscoped().Delete(&um)
	}
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user ID.
func DeleteUser(id uint64) error {
	if id == 1 {
		return berror.ErrDeleteAdmin
	}
	return DB.Self.Where("id = ?", id).Delete(&User{}).Error
}

// UpdateUser updates a user account information.
func (u *User) UpdateUser() error {
	return DB.Self.Save(u).Error
}

// GetUser gets a user by the user name
// returns user model, deleted and error.
func GetUser(username string) (u *User, deleted bool, err error) {
	u = &User{}
	r1 := DB.Self.Where("username = ?", username).First(&u)

	// found record
	if err = r1.Error; err == nil {
		return u, false, nil
	}
	r2 := DB.Self.Unscoped().Where("username = ?", username).First(&u)
	// not found record
	if errors.Is(r2.Error, gorm.ErrRecordNotFound) {
		return u, false, gorm.ErrRecordNotFound
	}
	// found record but deleted
	if errors.Is(r1.Error, gorm.ErrRecordNotFound) && r2.Error == nil {
		return u, true, nil
	}
	return u, false, nil
}

// GetUserByID gets a user model by ID
func GetUserByID(id uint64) (u *User, err error) {
	u = &User{}
	return u, DB.Self.Where("id = ?", id).First(&u).Error
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) error {
	return auth.Compare(u.Password, pwd)
}

// Encrypt the user password.
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) SetRole(role string) {
	u.Role = role
}

func (u *User) SetUserInfo(username string, password string) error {
	if username != "" {
		u.Username = username
	}
	if password != "" {
		u.Password = password
	}
	return nil
}
