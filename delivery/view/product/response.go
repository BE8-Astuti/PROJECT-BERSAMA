package product

import (
	"net/http"
)

type RespondProduct struct {
	ProductID   uint   `json:"productId"`
	UserID      uint   `json:"userId"`
	CategoryID  uint   `json:"categoryId"`
	Name        string `json:"name"`
	NameSeller  string `json:"nameSeller"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Sold        int    `json:"sold" gorm:"default:0"`
	UrlProduct  string `json:"urlProduct"`
}

func StatusGetAllOk(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All Data",
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

func StatusCreate(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Product",
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
