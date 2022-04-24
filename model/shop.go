package model

type ShopModel struct {
	BaseModel
	ShopName string       `json:"ShopName" gorm:"column:shopName;not null;" binding:"required" validation:"min=6,max=32"`
	Owner    *UserModel   `json:"owner" gorm:"column:owner;not null" binding:"required"`
	Staff    []*UserModel `json:"staff" gorm:"column:staff" binding:"required"`
	Phone    string       `json:"phone" gorm:"column:phone;not null" binding:"required" validate:"min=6,max=32"`
}

//TableName returns the table name.
func (s *ShopModel) TableName() string {
	return "tb_shops"
}

// CreteShop creates a shop model.
func (s *ShopModel) CreteShop(deleted bool) error {
	if deleted == true {
		sm := &ShopModel{}
		DB.Self.Unscoped().Where("shopName = ?", s.ShopName).First(&sm)
		DB.Self.Unscoped().Delete(&sm)
	}
	return DB.Self.Create(&s).Error
}
