package model

import . "Jinshuzhai-Bookstore/database"

type ShopBaseModel struct {
	Address     string `json:"address" gorm:"column:address"  validation:"min=5,max=128"`
	Phone       string `json:"phone" gorm:"column:phone" validation:"min=5,max=32"`
	Email       string `json:"email" gorm:"column:email"`
	Website     string `json:"website" gorm:"column:website" validation:"url"`
	Logo        string `json:"logo" gorm:"column:logo"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	IsOpen      bool   `json:"isOpen" gorm:"column:isOpen;default:false"`
}

type ShopModel struct {
	BaseModel
	//ShopBaseModel
	Name string `json:"name" gorm:"column:name;not null" binding:"required" validation:"min=5,max=32"`
	/*	OwnerID uint         `json:"ownerId" gorm:"column:ownerId;not null" binding:"required"`
		Owner   *UserModel   `json:"owner" gorm:"column:owner;not null" binding:"required"`
		Staff   []*UserModel `json:"staff" gorm:"column:staff" binding:"required"`*/
	Owner UserModel `json:"owner" gorm:"foreignkey:UserID" binding:"required"`
	//Staff []UserModel `json:"staff" gorm:"foreignkey:ShopID" binding:"required"`
}

//TableName returns the table name.
func (s *ShopModel) TableName() string {
	return "tb_shops"
}

// CreteShop creates a shop model.
func (s *ShopModel) CreteShop(deleted bool) error {
	if deleted == true {
		sm := &ShopModel{}
		DB.Self.Unscoped().Where("shopName = ?", s.Name).First(&sm)
		DB.Self.Unscoped().Delete(&sm)
	}
	return DB.Self.Create(&s).Error
}
