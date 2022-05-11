package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name    string    `json:"name"`
	Product []Product `gorm:"foreignKey:CategoryID;references:id"`
}
