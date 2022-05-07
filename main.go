package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `gorm:"unique" json:"password" form:"password"`

	Phone  string `json:"phone"`
	Status string `json:"status"`
	// Address []Address `gorm:"foreignKey:Owner"`
}
