package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
<<<<<<< HEAD
	Name        string        `json:"name"`
	Email       string        `json:"email" gorm:"unique"`
	Password    string        `json:"password" form:"password"`
	Phone       string        `json:"phone" gorm:"unique"`
	Status      bool          `json:"status"`
	Address     []Address     `gorm:"foreignKey:UserID;references:id"`
	Product     []Product     `gorm:"foreignKey:UserID;references:id"`
	Cart        []Cart        `gorm:"foreignKey:UserID;references:id"`
	Transaction []Transaction `gorm:"foreignKey:UserID;references:"`
=======
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"password" form:"password"`
	Phone    string    `json:"phone" gorm:"unique"`
	Status   string    `json:"status"`
	Address  []Address `gorm:"foreignKey:UserID;references:id"`
	Product  []Product `gorm:"foreignKey:UserID;references:id"`
	Cart     []Cart    `gorm:"foreignKey:UserID;references:id"`
>>>>>>> 198fdc5e30eb22b84568222bb16496948a9a28fd
}
