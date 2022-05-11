package routes

import (
	"together/be8/delivery/controller/address"
	"together/be8/delivery/controller/cart"
	"together/be8/delivery/controller/transaction"
	"together/be8/delivery/controller/user"

	// cbook "together/be8/delivery/controller/book"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Path(e *echo.Echo, u user.ControllerUser, a address.AddressControl, c cart.CartControl, t transaction.TransController) {

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// Login
	e.POST("/login", u.Login())
	// ROUTES USER
	user := e.Group("/user")
	user.POST("", u.InsertUser()) // Register
	// user.GET("", u.GetAllUser, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.GET("/:id", u.GetUserbyID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.PUT("/:id", u.UpdateUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.DELETE("/:id", u.DeleteUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))

	// Routes Addressy
	Address := e.Group("/address")
	Address.POST("", a.CreateAddress(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Address.GET("", a.GetAllAddress(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Address.GET("/:id", a.GetAddressID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Address.PUT("/:id", a.UpdateAddress(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Address.DELETE("/:id", a.DeleteAddress(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Address.PUT("/:id/default", a.SetDefaultAddress(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))

	// Routes Cart
	Cart := e.Group("/cart")
	Cart.POST("", c.CreateCart(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Cart.GET("", c.GetAllCart(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Cart.PUT("/:id", c.UpdateCart(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Cart.DELETE("/:id", c.DeleteCart(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Cart.GET("/shipment", c.Shipment(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))

	// ROUTES TRANSACTION
	Transaction := e.Group("/transaction")
	Transaction.POST("", t.CreateTransaction(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Transaction.GET("", t.GetAllTransaction(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Transaction.GET("/:order_id", t.GetTransactionDetail(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Transaction.POST("/:order_id/pay", t.PayTransaction(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	Transaction.POST("/:order_id/cancel", t.CancelTransaction(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
}
