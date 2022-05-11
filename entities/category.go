package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UserID int    `gorm:"user_id" json:"user_id"`
	Name   string `gorm:"name" json:"name"`
	// Product []Product `gorm:"foreignKey:CategoryID;references:id"`
}
