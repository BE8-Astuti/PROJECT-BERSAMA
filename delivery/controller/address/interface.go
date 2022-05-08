package address

import "github.com/labstack/echo/v4"

type AddressControl interface {
	CreateAddress() echo.HandlerFunc
	GetAllAddress() echo.HandlerFunc
	GetAddressID() echo.HandlerFunc
	UpdateAddress() echo.HandlerFunc
	DeleteAddress() echo.HandlerFunc
}
