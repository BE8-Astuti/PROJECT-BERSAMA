package product

import (
	"net/http"
	"together/be8/entities"
)

func StatusGetAllOk(data []entities.Product) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All data",
		"status":  true,
		"data":    data,
	}
}

func StatusGetIdOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Data",
		"status":  true,
		"data":    data,
	}
}

func StatusCreate(data entities.Product) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Product",
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(data entities.Product) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}
