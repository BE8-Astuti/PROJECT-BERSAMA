package user

import (
	"github.com/labstack/echo/v4"
)

type ControllerUser interface {
	InsertUser(c echo.Context) error // GetAllUser() echo.HandlerFunc
	GetUserbyID() echo.HandlerFunc
	UpdateUserID() echo.HandlerFunc
	DeleteUserID() echo.HandlerFunc
	Login() echo.HandlerFunc
}
