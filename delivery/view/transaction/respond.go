package transaction

import (
	"net/http"
	"together/be8/entities"
)

type AllTrans struct {
	TransDetail entities.Transaction
	Product     []entities.Cart
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

func StatusPayTrans(data entities.Transaction) map[string]interface{} {
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
