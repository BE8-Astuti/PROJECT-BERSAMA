package main

import (
	"fmt"
	"together/be8/config"
	cAddress "together/be8/delivery/controller/address"
	"together/be8/delivery/routes"
	"together/be8/repository/address"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	// Get Access Database
	database := config.InitDB()

	// Send AccessDB
	AddressRepo := address.NewDB(database)
	AddressControl := cAddress.NewControlAddress(AddressRepo, validator.New())
	fmt.Println(AddressControl)

	// Initiate Echo
	e := echo.New()
	// Akses Path Address
	routes.Path(e, AddressControl)
	e.Logger.Fatal(e.Start(":8000"))
}
