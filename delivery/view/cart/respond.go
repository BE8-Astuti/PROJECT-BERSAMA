package cart

import (
	"net/http"
	"together/be8/delivery/view/address"
)

type RespondCart struct {
	CartID      uint   `json:"cartId"`
	ProductID   uint   `json:"productId"`
	NameSeller  string `json:"nameSeller"`
	NameProduct string `json:"nameProduct"`
	Qty         int    `json:"qty"`
	Price       int    `json:"price"`
	ToBuy       string `json:"toBuy"`
	UrlProduct  string `json:"urlProduct"`
}

type CartProduct struct {
	CartID      uint   `json:"cartId"`
	ProductID   uint   `json:"productId"`
	NameProduct string `json:"nameProduct"`
	Qty         int    `json:"qty"`
	Price       int    `json:"price"`
	ToBuy       string `json:"toBuy"`
	UrlProduct  string `json:"urlProduct"`
}

type GetCart struct {
	NameSeller string        `json:"nameSeller"`
	Product    []CartProduct `json:"product"`
	SubTotal   int           `json:"subTotal"`
}

type Shipment struct {
	Address   address.RespondAddress `json:"address"`
	Product   []GetCart
	BillTotal int `json:"billTotal"`
}

func StatusGetAllOk(data []GetCart, BillTotal int) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All data",
		"status":  true,
		"data":    data,
		"bill":    BillTotal,
	}
}

func StatusGetIdOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Data ID",
		"status":  true,
		"data":    data,
	}
}

func StatusCreate(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Cart",
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}

func ShipmentOk(data Shipment) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Data Shipment",
		"status":  true,
		"data":    data,
	}
}
