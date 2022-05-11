package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UserID int    `gorm:"user_id" json:"user_id" validate:"required"`
	Name   string `gorm:"name" json:"name" validate:"required"`
}
