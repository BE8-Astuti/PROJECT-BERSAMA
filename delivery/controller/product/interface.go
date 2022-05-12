package product

import "github.com/labstack/echo/v4"

type ProductControl interface {
	InsertProd() echo.HandlerFunc
	GetProID() echo.HandlerFunc
	GetProdukbySeller() echo.HandlerFunc
	GetProdukByCategory() echo.HandlerFunc
	UpdateProduk() echo.HandlerFunc
	DeleteProduk() echo.HandlerFunc
	GetAllProduct() echo.HandlerFunc
}
