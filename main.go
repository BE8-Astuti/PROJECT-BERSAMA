package main

import (
	"together/be8/config"
	cAddress "together/be8/delivery/controller/address"
	cCart "together/be8/delivery/controller/cart"
	controllercat "together/be8/delivery/controller/category"
	catRepo "together/be8/repository/category"

	"together/be8/delivery/routes"
	"together/be8/repository/address"
	"together/be8/repository/cart"

	controllerprod "together/be8/delivery/controller/product"
	cTrans "together/be8/delivery/controller/transaction"

	produkRepo "together/be8/repository/product"
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

	productRepo := produkRepo.New(database)
	productControl := controllerprod.New(*productRepo, validator.New())

	catRepo := catRepo.NewDB(database)
	categoryControl := controllercat.NewControlCategory(catRepo, validator.New())
	// Initiate Echo

	// Send Access DB to Transaction
	snap := utils.InitMidtrans()
	transRepo := transaction.NewTransDB(database)
	transControl := cTrans.NewRepoTrans(transRepo, validator.New(), snap)

	// Initiate Echo
	e := echo.New()
	// Akses Path Addressss
	routes.Path(e, userControl, AddressControl, cartControl, transControl, categoryControl, productControl)
	config.Migrate()
	e.Logger.Fatal(e.Start(":8000"))
}
