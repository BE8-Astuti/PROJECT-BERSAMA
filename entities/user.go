package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string        `json:"username" gorm:"unique"`
	Name        string        `json:"name"`
	Email       string        `json:"email" gorm:"unique"`
	Password    string        `json:"password" form:"password"`
	Phone       string        `json:"phone" gorm:"unique"`
	Status      string        `json:"status"`
	Address     []Address     `gorm:"foreignKey:UserID;references:id"`
	Product     []Product     `gorm:"foreignKey:UserID;references:id"`
	Cart        []Cart        `gorm:"foreignKey:UserID;references:id"`
	Transaction []Transaction `gorm:"foreignKey:UserID;references:id"`
}
