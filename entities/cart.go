package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	NameSeller  string `json:"nameSeller"`
	NameProduct string `json:"nameProduct"`
	Qty         int    `json:"qty"`
	Price       int    `json:"price"`
}
