package product

import "together/be8/entities"

type RepoProduk interface {
	InsertProduk(newProduk entities.Product) (entities.Product, error)
	GetProdbyID(id uint) (entities.Product, error)
	GetProdBySeller(UserID uint) ([]entities.Product, error)
	GetProdByCategory(id int) ([]entities.Product, error)
	UpdateProduk(id int, UpdateProduk entities.Product, UserID uint) (entities.Product, error)
	DeleteProduk(id uint, UserID uint) error
	GetAllProduct() ([]entities.Product, error)
}
