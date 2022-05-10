package user

import (
	"github.com/labstack/echo/v4"
)

type ControllerUser interface {
<<<<<<< HEAD
	InsertUser(c echo.Context) error // GetAllUser() echo.HandlerFunc
=======
	InsertUser() echo.HandlerFunc
	// GetAllUser(c echo.Context) error
>>>>>>> 198fdc5e30eb22b84568222bb16496948a9a28fd
	GetUserbyID() echo.HandlerFunc
	UpdateUserID() echo.HandlerFunc
	DeleteUserID() echo.HandlerFunc
	Login() echo.HandlerFunc
}
