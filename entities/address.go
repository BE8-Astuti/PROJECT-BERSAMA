package entities

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	// UserID       uint
	Recipient    string `gorm:"recipient"`
	HP           string `gorm:"hp"`
	Street       string `gorm:"street"`
	SubDistrict  string `gorm:"subDistrict"`
	UrbanVillage string `gorm:"urbanVillage"`
	City         string `gorm:"city"`
	Zip          string `gorm:"zip"`
}
