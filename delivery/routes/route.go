package routes

import (
	"together/be8/delivery/controller/address"
	"together/be8/delivery/controller/cart"
	"together/be8/delivery/controller/category"
	"together/be8/delivery/controller/product"
	"together/be8/delivery/controller/transaction"
	"together/be8/delivery/controller/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Path(e *echo.Echo, u user.ControllerUser, a address.AddressControl, c cart.CartControl, t transaction.TransController, cat category.CategoryControl, p product.ProductControl) {

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// Login
	e.POST("/login", u.Login())
	// ROUTES USER
	user := e.Group("/user")
	user.POST("", u.InsertUser()) // Register
	// user.GET("", u.GetAllUser, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.GET("/:id", u.GetUserbyID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.PUT("/:id", u.UpdateUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	user.DELETE("/:id", u.DeleteUserID(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))

	category := e.Group("/category")
	category.PUT("/:id", cat.UpdateCat(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	category.DELETE("/:id", cat.DeleteCat(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	category.POST("", cat.CreateCategory(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	category.GET("", cat.GetAllCategory())
	category.GET("/:id", cat.GetCategoryID())

	product := e.Group("/product")
	product.POST("", p.InsertProd(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	product.GET("", p.GetAllProduct())
	product.GET("/:id", p.GetProID())
	product.PUT("/:id", p.UpdateProduk(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	product.DELETE("/:id", p.DeleteProduk(), middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("TOGETHER")}))
	product.GET("/user/:id", p.GetProdukbySeller())
	product.GET("/category/:id", p.GetProdukByCategory())

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
	Transaction.GET("/finish_payment", t.FinishPayment())
}
