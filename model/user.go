package model

import (
	. "Jinshuzhai-Bookstore/database"
	"Jinshuzhai-Bookstore/pkg/auth"
	"Jinshuzhai-Bookstore/pkg/constvar"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UserBaseModel represents base user information.
type UserBaseModel struct {
	Phone   string `json:"phone"   gorm:"column:phone;default:"   validate:"min=5,max=32"`
	Address string `json:"address" gorm:"column:address;default:" validate:"min=5,max=128"`
}

// UserModel represents user information.
type UserModel struct {
	BaseModel
	//UserBaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=2,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Role     string `json:"role"     gorm:"column:role;not null;default:general"        validate:"oneof=general seller admin"`
}

// TableName returns the table name.
func (u *UserModel) TableName() string {
	return "tb_users"
}

// CreateUser creates a new user account.
func (u *UserModel) CreateUser(deleted bool) error {
	if deleted == true {
		um := &UserModel{}
		DB.Self.Unscoped().Where("username = ?", u.Username).First(&um)
		DB.Self.Unscoped().Delete(&um)
	}
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user ID.
func DeleteUser(id uint64) error {
	if id == 1 {
		zap.L().Error("Tried to delete the admin user.")
		return errors.New("can not delete the admin user")
	}
	return DB.Self.Where("id = ?", id).Delete(&UserModel{}).Error
}

// UpdateUser updates a user account information.
func (u *UserModel) UpdateUser() error {
	return DB.Self.Save(u).Error
}

// GetUser gets a user by the user name
// returns user model, deleted and error.
func GetUser(username string) (um *UserModel, deleted bool, err error) {
	um = &UserModel{}
	d1 := DB.Self.Where("username = ?", username).First(&um)

	// found record
	if err := d1.Error; err == nil {
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
func GetUserByID(id uint64) (um *UserModel, err error) {
	um = &UserModel{}
	return um, DB.Self.Where("id = ?", id).First(&um).Error
}

// ListUser lists all users.
func ListUser(username string, pageNum, pageSize int) ([]*UserModel, int64, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	offset := (pageNum - 1) * pageSize

	userList := make([]*UserModel, 0)
	var count int64
	if len(username) > 0 {
		DB.Self.Where("username like ?", "%"+username+"%").Count(&count)
		if err := DB.Self.Where("username like ?", "%"+username+"%").Offset(offset).Limit(pageSize).Find(&userList).Error; err != nil {
			return userList, count, err
		}
	} else {
		DB.Self.Model(&UserModel{}).Count(&count)
		if err := DB.Self.Offset(offset).Limit(pageSize).Find(&userList).Error; err != nil {
			return userList, count, err
		}
	}
	return userList, count, nil
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
