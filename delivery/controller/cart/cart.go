package cart

import (
	"fmt"
	"net/http"
	"strconv"
	middlewares "together/be8/delivery/middleware"
	"together/be8/delivery/view"
	"together/be8/delivery/view/address"
	cartV "together/be8/delivery/view/cart"

	"together/be8/entities"
	"together/be8/repository/cart"

	"github.com/labstack/gommon/log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ControlCart struct {
	Repo  cart.CartRepository
	Valid *validator.Validate
}

func NewControlCart(NewCart cart.CartRepository, validate *validator.Validate) *ControlCart {
	return &ControlCart{
		Repo:  NewCart,
		Valid: validate,
	}
}

// ADD NEW CART
func (r *ControlCart) CreateCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Insert cartV.InsertCart
		if err := c.Bind(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}

		if err := r.Valid.Struct(&Insert); err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.Validate())
		}
		UserID := middlewares.ExtractTokenUserId(c)
		NewAdd := entities.Cart{
			UserID:    uint(UserID),
			ProductID: Insert.ProductID,
			Qty:       Insert.Qty,
		}
		result, errCreate := r.Repo.CreateCart(NewAdd)
		if errCreate != nil {
			log.Warn(errCreate)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		response := cartV.RespondCart{
			CartID:      result.ID,
			ProductID:   result.ProductID,
			NameSeller:  result.NameSeller,
			NameProduct: result.NameProduct,
			Qty:         result.Qty,
			Price:       result.Price,
			ToBuy:       result.ToBuy,
		}
		return c.JSON(http.StatusCreated, cartV.StatusCreate(response))
	}
}

// METHOD GET ALL CART
func (r *ControlCart) GetAllCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		result, seller, err := r.Repo.GetAllCart(uint(UserID))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		var data cartV.GetCart
		var res []cartV.GetCart
		var totalbill int
		cek := map[uint]int{}
		for _, NameSeller := range seller {
			var addProduct []cartV.CartProduct
			subTotal := 0
			for _, v := range result {
				if v.NameSeller == NameSeller {
					product := cartV.CartProduct{CartID: v.ID, ProductID: v.ProductID, NameProduct: v.NameProduct, Qty: v.Qty, Price: v.Price, ToBuy: v.ToBuy}
					addProduct = append(addProduct, product)
					if _, ok := cek[v.ID]; !ok {
						if v.ToBuy == "yes" {
							subTotal += v.Price * v.Qty
							fmt.Println(subTotal, NameSeller)
						}
						cek[v.ID]++
					}
				}
			}
			totalbill += subTotal
			data.Product = addProduct
			data.NameSeller = NameSeller
			data.SubTotal = subTotal
			res = append(res, data)
		}
		return c.JSON(http.StatusOK, cartV.StatusGetAllOk(res, totalbill))
	}
}

// UPDATE CART BY ID
func (r *ControlCart) UpdateCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var update cartV.UpdateCart
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}
		id := c.Param("id")
		idcart, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		UpdateCart := entities.Cart{Qty: update.Qty, ToBuy: update.ToBuy}
		UserID := middlewares.ExtractTokenUserId(c)
		result, errNotFound := r.Repo.UpdateCart(uint(idcart), UpdateCart, uint(UserID))
		if errNotFound != nil {
			log.Warn(errNotFound)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}

		response := cartV.RespondCart{
			CartID:      result.ID,
			ProductID:   result.ProductID,
			NameSeller:  result.NameSeller,
			NameProduct: result.NameProduct,
			Qty:         result.Qty,
			Price:       result.Price,
			ToBuy:       result.ToBuy,
		}

		return c.JSON(http.StatusOK, cartV.StatusUpdate(response))
	}
}

// DELETE CART BY ID
func (r *ControlCart) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idcart, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}

		UserID := middlewares.ExtractTokenUserId(c)

		errDelete := r.Repo.DeleteCart(uint(idcart), uint(UserID))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}

// GET SHIPMENT DATA
func (r *ControlCart) Shipment() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserId(c)
		Addr, Cart, seller, err := r.Repo.Shipment(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}

		AddressRes := address.RespondAddress{
			AddressID:      Addr.ID,
			Recipient:      Addr.Recipient,
			HP:             Addr.HP,
			Street:         Addr.Street,
			SubDistrict:    Addr.SubDistrict,
			UrbanVillage:   Addr.UrbanVillage,
			City:           Addr.City,
			Zip:            Addr.Zip,
			AddressDefault: Addr.AddressDefault,
		}

		var data cartV.GetCart
		var res []cartV.GetCart
		var totalbill int
		cek := map[uint]int{}
		for _, NameSeller := range seller {
			var addProduct []cartV.CartProduct
			subTotal := 0
			for _, v := range Cart {
				if v.NameSeller == NameSeller {
					product := cartV.CartProduct{CartID: v.ID, ProductID: v.ProductID, NameProduct: v.NameProduct, Qty: v.Qty, Price: v.Price, ToBuy: v.ToBuy}
					addProduct = append(addProduct, product)
					if _, ok := cek[v.ID]; !ok {
						if v.ToBuy == "yes" {
							subTotal += v.Price * v.Qty
							fmt.Println(subTotal, NameSeller)
						}
						cek[v.ID]++
					}
				}
			}
			totalbill += subTotal
			data.Product = addProduct
			data.NameSeller = NameSeller
			data.SubTotal = subTotal
			res = append(res, data)
		}
		Shipment := cartV.Shipment{Address: AddressRes, Product: res, BillTotal: totalbill}
		return c.JSON(http.StatusOK, cartV.ShipmentOk(Shipment))
	}
}
