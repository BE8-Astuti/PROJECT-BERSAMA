package routes

import (
	"together/be8/delivery/controller/address"

	"github.com/labstack/echo/v4"
)

func Path(e *echo.Echo, a address.AddressControl) {
	e.POST("/address", a.CreateAddress())
	e.GET("/address", a.GetAllAddress())
	e.GET("/address/:id", a.GetAddressID())
	e.PUT("/address/:id", a.UpdateAddress())
	e.DELETE("/address/:id", a.DeleteAddress())
}
