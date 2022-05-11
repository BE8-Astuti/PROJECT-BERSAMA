package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UserID  uint      `json:"userid"`
	Name    string    `json:"name"`
	Product []Product `gorm:"foreignKey:category_id"`
}
