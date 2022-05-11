package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	NameSeller  string `json:"nameSeller"`
	NameProduct string `json:"nameProduct"`
	Qty         int    `json:"qty"`
	Price       int    `json:"price"`
<<<<<<< HEAD
	ToBuy       string `json:"toBuy"`
=======
	ToBuy       string `json:"toBuy" gorm:"default:yes"`
	OrderID     string `json:"orderID"`
>>>>>>> fbedbae8ed32763c12abe5f92c0cbd8da656f0dc
	UserID      uint
	ProductID   uint
}
