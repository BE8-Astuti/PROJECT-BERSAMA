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

func StatusGetOrderOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Transaction Detail",
		"status":  true,
		"data":    data,
	}
}

func StatusCreate(data entities.Transaction, snap map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":        http.StatusCreated,
		"message":     "Success Create Transaction",
		"status":      true,
		"data":        data,
		"RedirectUrl": snap,
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
