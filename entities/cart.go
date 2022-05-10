package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	NameSeller  string `json:"nameSeller"`
	NameProduct string `json:"nameProduct"`
	Qty         int    `json:"qty"`
	Price       int    `json:"price"`
	ToBuy       string `json:"toBuy" gorm:"default:yes"`
	OrderID     string `json:"orderID"`
	UserID      uint
	ProductID   uint
}
