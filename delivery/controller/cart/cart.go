package cart

import (
	"fmt"
	"net/http"
	"strconv"
	"together/be8/delivery/view"
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

// METHOD Add New cart
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
		NewAdd := entities.Cart{
			// UserID:       1,
			NameSeller:  Insert.NameSeller,
			NameProduct: Insert.NameProduct,
			Qty:         Insert.Qty,
			Price:       Insert.Price,
			ToBuy:       Insert.ToBuy,
		}
		result, errCreate := r.Repo.CreateCart(NewAdd)
		if errCreate != nil {
			log.Warn(errCreate)
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusCreated, cartV.StatusCreate(result))
	}
}

// METHOD GET ALL cart
func (r *ControlCart) GetAllCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, seller, err := r.Repo.GetAllCart()
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
			for _, v := range result {
				if v.NameSeller == NameSeller {
					product := cartV.CartProduct{NameProduct: v.NameProduct, Qty: v.Qty, Price: v.Price, ToBuy: v.ToBuy}
					addProduct = append(addProduct, product)
				}
				if _, ok := cek[v.ID]; !ok {
					if v.ToBuy == true {
						bill := v.Price * v.Qty
						totalbill += bill
					}
					fmt.Println(totalbill, v.ID, cek)
					cek[v.ID]++
				}
			}
			data.Product = addProduct
			data.NameSeller = NameSeller
			res = append(res, data)
			fmt.Println(res)
		}

		return c.JSON(http.StatusOK, cartV.StatusGetAllOk(res, totalbill))
	}
}

// METHOD GET cart BY ID
func (r *ControlCart) GetCartID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idcart, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		result, errGetcartID := r.Repo.GetCartID(uint(idcart))
		if errGetcartID != nil {
			log.Warn(errGetcartID)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, cartV.StatusGetIdOk(result))
	}
}

// UPDATE cart BY ID
func (r *ControlCart) UpdateCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var update cartV.InsertCart
		if err := c.Bind(&update); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, view.BindData())
		}
		id := c.Param("id")
		idcart, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		Updatecart := entities.Cart{}

		if update.NameSeller != "" {
			Updatecart.NameSeller = update.NameSeller
		}
		if update.NameProduct != "" {
			Updatecart.NameProduct = update.NameProduct
		}
		if update.Qty != 0 {
			Updatecart.Qty = update.Qty
		}
		if update.Price != 0 {
			Updatecart.Price = update.Price
		}

		result, errNotFound := r.Repo.UpdateCart(uint(idcart), Updatecart)
		if errNotFound != nil {
			log.Warn(errNotFound)
			return c.JSON(http.StatusNotFound, view.NotFound())
		}
		return c.JSON(http.StatusOK, cartV.StatusUpdate(result))
	}
}

// DELETE cart BY ID
func (r *ControlCart) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idcart, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, view.ConvertID())
		}
		errDelete := r.Repo.DeleteCart(uint(idcart))
		if errDelete != nil {
			return c.JSON(http.StatusInternalServerError, view.InternalServerError())
		}
		return c.JSON(http.StatusOK, view.StatusDelete())
	}
}
