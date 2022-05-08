package main

import (
	"fmt"
	"log"
	"together/be8/config"
	cAddress "together/be8/delivery/controller/address"
	cCart "together/be8/delivery/controller/cart"
	"together/be8/delivery/routes"
	"together/be8/repository/address"
	"together/be8/repository/cart"

	controllerus "together/be8/delivery/controller/user"
	"together/be8/delivery/routes"
	userRepo "together/be8/repository/user"

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
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	config.Migrate(db)
	e := echo.New()

	repo := userRepo.New(db)
	controller := controllerus.New(repo, validator.New())
	routes.RegisterPath(e, controller)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
