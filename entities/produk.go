package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID      uint
	CategoryID  uint
	Name        string `json:"name"`
	NameSeller  string `json:"nameSeller"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Sold        int    `json:"sold" gorm:"default:0"`
	UrlProduct  string `json:"urlProduct"`
}
