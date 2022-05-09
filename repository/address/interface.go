package address

import "together/be8/entities"

type RepoAddress interface {
	CreateAddress(newAdd entities.Address) (entities.Address, error)
	GetAllAddress(UserID uint) ([]entities.Address, error)
	GetAddressID(id uint, UserID uint) (entities.Address, error)
	UpdateAddress(id uint, updatedAddress entities.Address, UserID uint) (entities.Address, error)
	DeleteAddress(id uint, UserID uint) error
	SetDefaultAddress(id uint, UserID uint) error
}
