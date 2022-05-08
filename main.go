package main

import (
	"together/be8/config"
	cAddress "together/be8/delivery/controller/address"
	cCart "together/be8/delivery/controller/cart"
	"together/be8/delivery/routes"
	"together/be8/repository/address"
	"together/be8/repository/cart"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	// Get Access Database
	database := config.InitDB()

	// Send AccessDB to AddressRepo
	AddressRepo := address.NewDB(database)
	AddressControl := cAddress.NewControlAddress(AddressRepo, validator.New())

	// Send Access DB to CartRepo
	cartRepo := cart.NewRepoCart(database)
	cartControl := cCart.NewControlCart(cartRepo, validator.New())

	// Initiate Echo
	e := echo.New()
	// Akses Path Address
	routes.Path(e, AddressControl, cartControl)
	e.Logger.Fatal(e.Start(":8000"))
}
