package address

import (
	"errors"
	"together/be8/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AddressDB struct {
	Db *gorm.DB
}

// Get Access DB
func NewDB(db *gorm.DB) *AddressDB {
	return &AddressDB{
		Db: db,
	}
}

// CREATE NEW ADDRESS TO DATABASE
func (a *AddressDB) CreateAddress(newAdd entities.Address) (entities.Address, error) {
	if err := a.Db.Create(&newAdd).Error; err != nil {
		log.Warn(err)
		return newAdd, err
	}
	return newAdd, nil
}

// GET ALL ADDRESS IN DATABASE
func (a *AddressDB) GetAllAddress(UserID uint) ([]entities.Address, error) {
	var AllAddress []entities.Address
	if err := a.Db.Where("user_id=?", UserID).Find(&AllAddress).Error; err != nil {
		log.Warn("Error Get All Address", err)
		return AllAddress, errors.New("Access Database Error")
	}

	return AllAddress, nil
}

// GET ADDRESS BY ID
func (a *AddressDB) GetAddressID(id uint, UserID uint) (entities.Address, error) {
	var Address entities.Address
	if err := a.Db.Where("id = ? AND user_id = ?", id, UserID).First(&Address).Error; err != nil {
		log.Warn("Error Get Address By ID", err)
		return Address, errors.New("Access Database Error")
	}
	return Address, nil
}

// UPDATE ADDRESS BY ID
func (a *AddressDB) UpdateAddress(id uint, updatedAddress entities.Address, UserID uint) (entities.Address, error) {
	var updated entities.Address

	if err := a.Db.Where("id =? AND user_id =?", id, UserID).First(&updated).Updates(&updatedAddress).Find(&updated).Error; err != nil {
		log.Warn("Update Address Error", err)
		return updated, errors.New("Access Database Error")
	}

	return updated, nil
}

// DELETE ADDRESS BY ID
func (a *AddressDB) DeleteAddress(id uint, UserID uint) error {

	var delete entities.Address
	if err := a.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	return nil
}

func (a *AddressDB) SetDefaultAddress(id uint, UserID uint) error {
	if err := a.Db.Table("addresses").Where("user_id = ?", UserID).Update("address_default", "no").Error; err != nil {
		return err
	}

	var address entities.Address

	if err := a.Db.Table("addresses").Where("id = ? AND user_id = ?", id, UserID).First(&address).Update("address_default", "yes").Error; err != nil {
		return err
	}
	return nil
}
