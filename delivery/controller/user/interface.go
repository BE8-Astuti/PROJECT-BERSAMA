package user

import (
	"github.com/labstack/echo/v4"
)

type ControllerUser interface {
	InsertUser(c echo.Context) error
	GetAllUser(c echo.Context) error
	GetUserbyID(c echo.Context) error
	UpdateUserID(c echo.Context) error
	DeleteUserID(c echo.Context) error
	Login(c echo.Context) error
}
