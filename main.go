package main

import (
	"together/be8/config"
	cAddress "together/be8/delivery/controller/address"
	cCart "together/be8/delivery/controller/cart"
	cTrans "together/be8/delivery/controller/transaction"
	"together/be8/delivery/routes"
	"together/be8/repository/address"
	"together/be8/repository/cart"
	"together/be8/repository/transaction"
	"together/be8/utils"

	controllerus "together/be8/delivery/controller/user"
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

	userRepo := userRepo.New(database)
	userControl := controllerus.New(userRepo, validator.New())

	// Send Access DB to Transaction
	snap := utils.InitMidtrans()
	transRepo := transaction.NewTransDB(database)
	transControl := cTrans.NewRepoTrans(transRepo, validator.New(), snap)

	// Initiate Echo
	e := echo.New()
	// Akses Path Addressss
	routes.Path(e, userControl, AddressControl, cartControl, transControl)
	e.Logger.Fatal(e.Start(":8000"))
}
