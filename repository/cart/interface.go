package cart

import "together/be8/entities"

type CartRepository interface {
	CreateCart(newAdd entities.Cart) (entities.Cart, error)
	GetAllCart() ([]entities.Cart, []string, error)
	GetCartID(id uint) (entities.Cart, error)
	UpdateCart(id uint, updatedCart entities.Cart) (entities.Cart, error)
	DeleteCart(id uint) error
}
