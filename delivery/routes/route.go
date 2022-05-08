package routes

import (
	"together/be8/delivery/controller/address"
	"together/be8/delivery/controller/cart"

	// cbook "together/be8/delivery/controller/book"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Path(e *echo.Echo,u user., a address.AddressControl, c cart.CartControl) {

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// ROUTES USER
	user := e.Group("/user")
	user.POST("/user", uc.InsertUser) // Register
	user.POST("/login", uc.Login)     // Login
	user.GET("/user", uc.GetAllUser, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.GET("/user/:id", uc.GetUserbyID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.PUT("/user/:id", uc.UpdateUserID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.DELETE("/user/:id", uc.DeleteUserID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))

	// Routes Address
	Address := e.Group("/address")
	Address.POST("", a.CreateAddress(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))
	Address.GET("", a.GetAllAddress(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))
	Address.GET("/:id", a.GetAddressID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))
	Address.PUT("/:id", a.UpdateAddress(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))
	Address.DELETE("/:id", a.DeleteAddress(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))

	// Routes Cart
	Cart := e.Group("/cart")
	Cart.POST("", c.CreateCart(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))
	Cart.GET("", c.GetAllCart(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))
	Cart.GET("/:id", c.GetCartID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))
	Cart.PUT("/:id", c.UpdateCart(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))
	Cart.DELETE("/:id", c.DeleteCart(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("K3YT0K3N")}))
}

