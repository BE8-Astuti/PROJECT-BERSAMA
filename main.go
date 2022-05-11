package main

import (
	"together/be8/config"
	cAddress "together/be8/delivery/controller/address"
	cCart "together/be8/delivery/controller/cart"
	controllercat "together/be8/delivery/controller/category"
	"together/be8/entities"
	catRepo "together/be8/repository/category"

	"together/be8/delivery/routes"
	"together/be8/repository/address"
	"together/be8/repository/cart"

	controllerprod "together/be8/delivery/controller/product"
	rprod "together/be8/repository/product"

	controllerus "together/be8/delivery/controller/user"
	userRepo "together/be8/repository/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	// Get Access Database

	database := config.InitDB()

	database.AutoMigrate(entities.Category{})
	database.AutoMigrate(entities.Address{})
	database.AutoMigrate(entities.Cart{})
	database.AutoMigrate(entities.Product{})
	// database.AutoMigrate(entities.Transaction{})

	// Send AccessDB to AddressRepo
	AddressRepo := address.NewDB(database)
	AddressControl := cAddress.NewControlAddress(AddressRepo, validator.New())

	// Send Access DB to CartRepo
	cartRepo := cart.NewRepoCart(database)
	cartControl := cCart.NewControlCart(cartRepo, validator.New())

	userRepo := userRepo.New(database)
	userControl := controllerus.New(userRepo, validator.New())

	prodrep := rprod.New(database)
	produkControl := controllerprod.New(prodrep, validator.New())

	catRepo := catRepo.NewDB(database)
	catControl := controllercat.NewControlCategory(catRepo, validator.New())
	// Initiate Echo
	e := echo.New()
	// Akses Path Addressss
	routes.Path(e, userControl, AddressControl, cartControl, catControl, produkControl)
	e.Logger.Fatal(e.Start(":8000"))

}
