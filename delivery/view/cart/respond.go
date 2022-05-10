package cart

import (
	"net/http"
	"together/be8/entities"
)

type CartProduct struct {
	NameProduct string `json:"nameProduct" validate:"required"`
	Qty         int    `json:"qty" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	ToBuy       string `json:"toBuy" validate:"required"`
}

type GetCart struct {
	NameSeller string        `json:"nameSeller"`
	Product    []CartProduct `json:"product"`
	SubTotal   int           `json:"subTotal"`
}

type Shipment struct {
	Address   entities.Address `json:"address"`
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

func StatusCreate(data entities.Cart) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Cart",
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(data entities.Cart) map[string]interface{} {
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
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}
