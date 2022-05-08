package routes

import (
	"together/be8/delivery/controller/address"
	"together/be8/delivery/controller/cart"

	// cbook "together/be8/delivery/controller/book"
	cuser "together/be8/delivery/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Path(e *echo.Echo, a address.AddressControl, c cart.CartControl) {

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// ROUTES USER
	user := e.Group("/user")
	e.POST("/user", uc.InsertUser) // Register
	e.POST("/login", uc.Login)     // Login
	e.GET("/user", uc.GetAllUser, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	e.GET("/user/:id", uc.GetUserbyID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	e.PUT("/user/:id", uc.UpdateUserID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	e.DELETE("/user/:id", uc.DeleteUserID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))

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
func RegisterPath(e *echo.Echo, uc cuser.ControllerUser) {
	// e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	// e.GET("/pegawai", pc.GetAllPegawai)
	// e.POST("/books", bc.InsertBook, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RH$SI4")}))
	// e.GET("/books", bc.GetAllBook)
	// e.GET("/books/:id", bc.GetBookbyID)
	// e.PUT("/books/:id", bc.UpdateBookID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RH$SI4")}))
	// e.DELETE("/user/:id", bc.DeleteBookID, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("RH$SI4")}))

}
