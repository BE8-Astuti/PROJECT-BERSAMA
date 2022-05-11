package address

type InsertAddress struct {
	Recipient    string `gorm:"recipient" json:"recipient" validate:"required"`
	HP           string `gorm:"hp" json:"hp" validate:"required"`
	Street       string `gorm:"street" json:"street" validate:"required"`
	SubDistrict  string `gorm:"subDistrict" json:"subDistrict" validate:"required"`
	UrbanVillage string `gorm:"urbanVillage" json:"urbanVillage" validate:"required"`
	City         string `gorm:"city" json:"city" validate:"required"`
	Zip          string `gorm:"zip" json:"zip" validate:"required"`
}
