package address

import "together/be8/entities"

type RepoAddress interface {
<<<<<<< HEAD
	CreateAddress(newAdd entities.Address) (entities.Address, error)
=======
	CreateAddress(newAdd entities.Address, UserID uint) (entities.Address, error)
>>>>>>> fbedbae8ed32763c12abe5f92c0cbd8da656f0dc
	GetAllAddress(UserID uint) ([]entities.Address, error)
	GetAddressID(id uint, UserID uint) (entities.Address, error)
	UpdateAddress(id uint, updatedAddress entities.Address, UserID uint) (entities.Address, error)
	DeleteAddress(id uint, UserID uint) error
	SetDefaultAddress(id uint, UserID uint) error
}
