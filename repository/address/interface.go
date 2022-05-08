package address

import "together/be8/entities"

type RepoAddress interface {
	CreateAddress(newAdd entities.Address) (entities.Address, error)
	GetAllAddress() ([]entities.Address, error)
	GetAddressID(id uint) (entities.Address, error)
	UpdateAddress(id uint, updatedAddress entities.Address) (entities.Address, error)
	DeleteAddress(id uint) error
}
