package main

import (
	"fmt"
	"log"
	"together/be8/config"

	controllerus "together/be8/delivery/controller/user"
	"together/be8/delivery/routes"
	userRepo "together/be8/repository/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	config.Migrate(db)
	e := echo.New()

	repo := userRepo.New(db)
	controller := controllerus.New(repo, validator.New())
	routes.RegisterPath(e, controller)
	log.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
