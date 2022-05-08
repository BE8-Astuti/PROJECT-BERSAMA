package user

import (
	"net/http"
	"together/be8/entities"
)

type LoginResponse struct {
	Token string
}

func SuccessInsert(data entities.User) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data user",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data user",
		"status":  false,
		"data":    nil,
	}
}

func LoginOK(data LoginResponse, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"status":  true,
		"data":    data,
	}
}
