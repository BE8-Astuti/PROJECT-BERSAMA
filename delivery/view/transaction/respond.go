package transaction

import (
	"net/http"
	"time"
)

type RespondTransaction struct {
	OrderID       string    `json:"orderID"`
	TotalBill     int       `json:"totalBill"`
	PaymentMethod string    `json:"paymentMethod"`
	Address       string    `json:"address"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
}

type ProductTransaction struct {
	ProductID   uint   `json:"productId"`
	NameSeller  string `json:"nameSeller"`
	NameProduct string `json:"nameProduct" validate:"required"`
	Qty         int    `json:"qty" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	UrlProduct  string `json:"urlProduct"`
	SubTotal    int    `json:"subTotal"`
}

type AllTrans struct {
	TransDetail RespondTransaction
	Product     []ProductTransaction
}

func StatusGetAllOk(data []AllTrans) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All data",
		"status":  true,
		"data":    data,
	}
}

func StatusTransactionDetail(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Transaction Detail",
		"status":  true,
		"data":    data,
	}
}

func StatusCreate(OrderID string, snap map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Transaction",
		"status":  true,
		"data":    map[string]interface{}{"order-id": OrderID, "RedirectUrl": snap},
	}
}

func StatusPayTrans(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Pay Transaction",
		"status":  true,
		"data":    data,
	}
}

func StatusCancelTrans() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Cancel Transaction",
		"status":  true,
	}
}

func StatusErrorSnap() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNoContent,
		"message": "Error Get Redirect Url Payment",
		"status":  false,
	}
}
