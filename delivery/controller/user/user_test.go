package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"together/be8/entities"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUser(t *testing.T) {
	t.Run("Success Get All", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&mockUserRepository{}, validator.New())
		userController.GetAllUser(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entities.User
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, resp.Data[0].Name, "Astuti")
	})
	t.Run("Error Get All", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.GetAllUser(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    []entities.User
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Nil(t, resp.Data)
		assert.False(t, resp.Status)
		assert.Equal(t, 500, resp.Code)
	})
}

func TestInsertUser(t *testing.T) {
	t.Run("Success Insert", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "yani",
			"email":    "y",
			"password": "849",
			"phone":    "77979799",
			"status":   false,
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set Content to JSON
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&mockUserRepository{}, validator.New())
		userController.InsertUser(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, "yani", resp.Data.(map[string]interface{})["name"])
		assert.True(t, resp.Status)
		assert.Equal(t, 201, resp.Code)
	})
	t.Run("Error Validasi", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"phone": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.InsertUser(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 400, resp.Code)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"phone": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.InsertUser(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		log.Warn(resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 400, resp.Code)
	})
	t.Run("Error Insert DB", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":  "yani",
			"phone": "779",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.InsertUser(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.False(t, resp.Status)
		assert.Nil(t, resp.Data)
		assert.Equal(t, 500, resp.Code)
	})
}

func TestGetUserbyID(t *testing.T) {
	t.Run("Success Get User by ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("6")

		userController := New(&mockUserRepository{}, validator.New())
		userController.GetUserbyID(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    entities.User
		}

		var resp response
		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, "yani", resp.Data.Name)
		assert.True(t, resp.Status)
		assert.Equal(t, 200, resp.Code)
	})
	t.Run("Error Konversi", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.GetUserbyID(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    entities.User
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Nil(t, resp.Data)
		assert.False(t, resp.Status)
		assert.Equal(t, 500, resp.Code)
	})
	t.Run("Error Get DB", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.GetUserbyID(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    entities.User
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Nil(t, resp.Data)
		assert.False(t, resp.Status)
		assert.Equal(t, 500, resp.Code)
	})
	t.Run("Error Get data", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user")

		userController := New(&erorrMockUserRepository{}, validator.New())
		userController.GetUserbyID(context)

		type response struct {
			Code    int
			Message string
			Status  bool
			Data    entities.User
		}

		var resp response

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Nil(t, resp.Data)
		assert.False(t, resp.Status)
		assert.Equal(t, 404, resp.Code)
	})
}

// func TestUpdateUserID(t *testing.T) {
// 	t.Run("Success Get All", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		userController := New(&mockUserRepository{}, validator.New())
// 		userController(context)

// 		type response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    []entity.Pegawai
// 		}

// 		var resp response

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Equal(t, resp.Data[0].Nama, "Jerry")
// 	})
// 	t.Run("Error Get All", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		pegawaiController := New(&erorrMockUserRepository{}, validator.New())
// 		pegawaiController.GetAllPegawai(context)

// 		type response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    []entity.Pegawai
// 		}

// 		var resp response

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Nil(t, resp.Data)
// 		assert.False(t, resp.Status)
// 		assert.Equal(t, 500, resp.Code)
// 	})
// }

// func TestDeleteUserID(t *testing.T) {
// 	t.Run("Success Get All", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		pegawaiController := New(&mockUserRepository{}, validator.New())
// 		pegawaiController.GetAllPegawai(context)

// 		type response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    []entity.Pegawai
// 		}

// 		var resp response

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Equal(t, resp.Data[0].Nama, "Jerry")
// 	})
// 	t.Run("Error Get All", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		pegawaiController := New(&erorrMockUserRepository{}, validator.New())
// 		pegawaiController.GetAllPegawai(context)

// 		type response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    []entity.Pegawai
// 		}

// 		var resp response

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Nil(t, resp.Data)
// 		assert.False(t, resp.Status)
// 		assert.Equal(t, 500, resp.Code)
// 	})
// }

// func TestLogin(t *testing.T) {
// 	t.Run("Success Get All", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		pegawaiController := New(&mockUserRepository{}, validator.New())
// 		pegawaiController.GetAllPegawai(context)

// 		type response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    []entity.Pegawai
// 		}

// 		var resp response

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Equal(t, resp.Data[0].Nama, "Jerry")
// 	})
// 	t.Run("Error Get All", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users")

// 		pegawaiController := New(&erorrMockUserRepository{}, validator.New())
// 		pegawaiController.GetAllPegawai(context)

// 		type response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    []entity.Pegawai
// 		}

// 		var resp response

// 		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
// 		assert.Nil(t, resp.Data)
// 		assert.False(t, resp.Status)
// 		assert.Equal(t, 500, resp.Code)
// 	})
// }

// Dummy Data

type mockUserRepository struct{}

func (mur *mockUserRepository) InsertUser(newUser entities.User) (entities.User, error) {
	return newUser, nil
}

func (mur *mockUserRepository) GetAllUser() ([]entities.User, error) {
	return []entities.User{{Name: "Astuti", Phone: "7897787", Email: "a@gmail.com", Status: false}}, nil
}

func (mur *mockUserRepository) GetUserID(ID int) (entities.User, error) {
	return entities.User{}, nil
}

func (mur *mockUserRepository) UpdateUser(ID int, email string) (entities.User, error) {
	return entities.User{}, nil
}

func (mur *mockUserRepository) DeleteUser(ID int) (entities.User, error) {
	return entities.User{}, nil
}

func (mur *mockUserRepository) Login(email, password string) (entities.User, error) {
	return entities.User{}, nil
}

type erorrMockUserRepository struct{}

func (emur *erorrMockUserRepository) InsertUser(newPegawai entities.User) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa insert data")
}
func (emur *erorrMockUserRepository) GetAllUser() ([]entities.User, error) {
	return nil, errors.New("tidak bisa select data")
}

func (emur *erorrMockUserRepository) DeleteUser(ID int) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa select data")
}

func (emur *erorrMockUserRepository) GetUserID(ID int) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa select data")
}

func (emur *erorrMockUserRepository) Login(email, password string) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa select data")
}
func (emur *erorrMockUserRepository) UpdateUser(ID int, email string) (entities.User, error) {
	return entities.User{}, errors.New("tidak bisa select data")
}
