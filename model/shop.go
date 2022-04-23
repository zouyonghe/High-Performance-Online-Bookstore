package model

type ShopModel struct {
	BaseModel
	Owner UserModel
	Staff []UserModel
	Phone string
}
