package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID      uint
<<<<<<< HEAD
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
=======
	CategoryID  uint
	Name        string `json:"name"`
	NameSeller  string `json:"nameSeller"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Sold        int    `json:"sold" gorm:"default:0"`
>>>>>>> fbedbae8ed32763c12abe5f92c0cbd8da656f0dc
}
