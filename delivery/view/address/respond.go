package address

import (
	"net/http"
)

type RespondAddress struct {
	AddressID      uint   `json:"addressId"`
	Recipient      string `json:"recipient"`
	HP             string `json:"hp"`
	Street         string `json:"street"`
	SubDistrict    string `json:"subDistrict"`
	UrbanVillage   string `json:"urbanVillage"`
	City           string `json:"city"`
	Zip            string `json:"zip"`
	AddressDefault string `json:"addressDefault"`
}

func StatusGetAllOk(data interface{}) map[string]interface{} {
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
		"message": "Success Get Data ID",
		"status":  true,
		"data":    data,
	}
}

func StatusCreate(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Success Create Address",
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

func StatusDefaultAddress() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Update Default Address Success",
		"status":  true,
	}
}
