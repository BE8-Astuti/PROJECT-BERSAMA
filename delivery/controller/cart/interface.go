package cart

import "github.com/labstack/echo/v4"

type CartControl interface {
	CreateCart() echo.HandlerFunc
	GetAllCart() echo.HandlerFunc
<<<<<<< HEAD
	GetCartID() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
=======
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
	Shipment() echo.HandlerFunc
>>>>>>> fbedbae8ed32763c12abe5f92c0cbd8da656f0dc
}
