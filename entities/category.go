package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
<<<<<<< HEAD
	Name string `gorm:"name" json:"name" validate:"required"`
=======
	Name string `json:"name"`
	Product []Product `gorm:"foreignKey:CategoryID;references:id"`
>>>>>>> fbedbae8ed32763c12abe5f92c0cbd8da656f0dc
}
