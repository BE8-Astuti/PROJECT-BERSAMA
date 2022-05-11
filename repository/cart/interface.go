package cart

import (
	"together/be8/entities"
)

type CartRepository interface {
	CreateCart(newAdd entities.Cart) (entities.Cart, error)
	GetAllCart(UserID uint) ([]entities.Cart, []string, error)
<<<<<<< HEAD
	GetCartID(id uint, UserID uint) (entities.Cart, error)
	UpdateCart(id uint, updatedCart entities.Cart, UserID uint) (entities.Cart, error)
	DeleteCart(id uint, UserID uint) error
=======
	UpdateCart(id uint, updatedCart entities.Cart, UserID uint) (entities.Cart, error)
	DeleteCart(id uint, UserID uint) error
	Shipment(UserID uint) (entities.Address, []entities.Cart, []string, error)
>>>>>>> fbedbae8ed32763c12abe5f92c0cbd8da656f0dc
}
