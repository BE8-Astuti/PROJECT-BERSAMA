package routes

import (
	// cbook "together/be8/delivery/controller/book"
	cuser "together/be8/delivery/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc cuser.ControllerUser) {
	// e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	// e.GET("/pegawai", pc.GetAllPegawai)
	e.POST("/user", uc.InsertUser) // Register
	e.POST("/login", uc.Login)     // Login
	e.GET("/user", uc.GetAllUser, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	e.GET("/user/:id", uc.GetUserbyID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	e.PUT("/user/:id", uc.UpdateUserID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	e.DELETE("/user/:id", uc.DeleteUserID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	// e.POST("/books", bc.InsertBook, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RH$SI4")}))
	// e.GET("/books", bc.GetAllBook)
	// e.GET("/books/:id", bc.GetBookbyID)
	// e.PUT("/books/:id", bc.UpdateBookID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RH$SI4")}))
	// e.DELETE("/user/:id", bc.DeleteBookID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RH$SI4")}))

}
