package model

import (
	"Jinshuzhai-Bookstore/pkg/auth"
	"Jinshuzhai-Bookstore/pkg/constvar"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

// UserModel represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Role     string `json:"role"     gorm:"column:role;not null;default:general"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

// Create creates a new user account.
func (u *UserModel) Create(deleted bool) error {
	if deleted == true {
		// zap.L().Info("delete called")
		um := &UserModel{}
		DB.Self.Unscoped().Where("username = ?", u.Username).First(&um)
		DB.Self.Unscoped().Delete(&um)
		//return DB.Self.Unscoped().Save(&u).Error
	}
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user ID.
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.ID = id
	return DB.Self.Delete(&user).Error
}

// Update updates a user account information.
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetUser gets a user by the user name
// return user model, deleted and error.
func GetUser(username string) (um *UserModel, deleted bool, err error) {
	um = &UserModel{}
	d1 := DB.Self.Where("username = ?", username).First(&um)

	// found record
	if d1.Error == nil {
		// zap.L().Info("found record called")
		return um, false, nil
	}
	d2 := DB.Self.Unscoped().Where("username = ?", username).First(&um)
	// not found record
	if errors.Is(d2.Error, gorm.ErrRecordNotFound) {
		return um, false, gorm.ErrRecordNotFound
	}
	// found record but deleted
	if errors.Is(d1.Error, gorm.ErrRecordNotFound) && d2.Error == nil {
		return um, true, nil
	}
	return um, false, nil
}

// GetUserByID gets a user model by ID
func GetUserByID(ID uint64) (um *UserModel, err error) {
	um = &UserModel{}
	d1 := DB.Self.Where("id = ?", ID).First(&um)

	// found record
	if d1.Error == nil {
		return um, nil
	}
	return um, d1.Error
}

// ListUser lists all users.
func ListUser(username string, offset, limit int) ([]*UserModel, int64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count int64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *UserModel) GetRole() string {
	return u.Role
}
