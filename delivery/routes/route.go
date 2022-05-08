package routes

import (
	"together/be8/delivery/controller/address"
	"together/be8/delivery/controller/cart"

	"github.com/labstack/echo/v4"
)

func Path(e *echo.Echo, a address.AddressControl, c cart.CartControl) {
	// Routes Address
	Address := e.Group("/address")
	Address.POST("", a.CreateAddress())
	Address.GET("", a.GetAllAddress())
	Address.GET("/:id", a.GetAddressID())
	Address.PUT("/:id", a.UpdateAddress())
	Address.DELETE("/:id", a.DeleteAddress())

	// Routes Cart
	Cart := e.Group("/cart")
	Cart.POST("", c.CreateCart())
	Cart.GET("", c.GetAllCart())
	Cart.GET("/:id", c.GetCartID())
	Cart.PUT("/:id", c.UpdateCart())
	Cart.DELETE("/:id", c.DeleteCart())
}
