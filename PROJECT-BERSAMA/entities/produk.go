package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID      uint
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
