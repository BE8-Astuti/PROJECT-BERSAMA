package cart

import "github.com/labstack/echo/v4"

type CartControl interface {
	CreateCart() echo.HandlerFunc
	GetAllCart() echo.HandlerFunc
	GetCartID() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
	Shipment() echo.HandlerFunc
}
