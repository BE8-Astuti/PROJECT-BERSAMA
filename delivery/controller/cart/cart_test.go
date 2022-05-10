package cart

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	middlewares "together/be8/delivery/middleware"
	"together/be8/entities"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var token string

// INITIATE TOKEN
func TestCreateToken(t *testing.T) {
	t.Run("Create Token", func(t *testing.T) {
		token, _ = middlewares.CreateToken(1, "Motor", "Motor@gmail.com")
	})
}

// TEST METHODE CREATE_ADDRESS
func TestCreateCart(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"productId":   3,
			"nameSeller":  "Elec Center",
			"nameProduct": "Mouse",
			"qty":         3,
			"price":       100000,
			"toBuy":       "yes",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart")
		CartC := NewControlCart(&mockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(CartC.CreateCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Create Cart", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"productId":   3,
			"nameSeller":  "Elec Center",
			"nameProduct": "Mouse",
			"qty":         3,
			"price":       100000,
			"toBuy":       "yes",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart")
		CartC := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(CartC.CreateCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody := "Jalan Gunung"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart")
		CartC := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(CartC.CreateCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Validate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"productId":   3,
			"nameSeller":  "Elec Center",
			"nameProduct": "Mouse",
			"qty":         3,
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart")
		CartC := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(CartC.CreateCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Validate Error", result.Message)
		assert.False(t, result.Status)
	})
}

// TEST GET ALL ADDRESS
func TestGetAllCart(t *testing.T) {
	t.Run("Success Get All Cart", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart")
		GetCart := NewControlCart(&mockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.GetAllCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart")
		GetCart := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.GetAllCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
}

// TEST GET ADDRESS BY ID
func TestGetCartID(t *testing.T) {
	t.Run("Success Get Cart By ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetCart := NewControlCart(&mockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.GetCartID())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response

		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get Data ID", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetCart := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.GetCartID())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetCart := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.GetCartID())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

// TEST UPDATE ADDRESS BY ID
func TestUpdateCart(t *testing.T) {
	t.Run("Update Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"qty":   7,
			"toBuy": "no",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		GetCart := NewControlCart(&mockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.UpdateCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Updated", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetCart := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.UpdateCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 404, result.Code)
		assert.Equal(t, "Data Not Found", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Bind", func(t *testing.T) {
		e := echo.New()
		requestBody := "Jalan Gunung"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		CartC := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(CartC.UpdateCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 415, result.Code)
		assert.Equal(t, "Cannot Bind Data", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetCart := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.UpdateCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

// TEST DELETE ADDRESS BY ID
func TestDeleteCart(t *testing.T) {
	t.Run("Success Delete Cart", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetCart := NewControlCart(&mockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.DeleteCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Deleted", result.Message)
		assert.True(t, result.Status)
	})
	t.Run("Error Delete Cart", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		GetCart := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.DeleteCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 500, result.Code)
		assert.Equal(t, "Cannot Access Database", result.Message)
		assert.False(t, result.Status)
	})
	t.Run("Error Convert ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/cart/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		GetCart := NewControlCart(&errMockCart{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetCart.DeleteCart())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 406, result.Code)
		assert.Equal(t, "Cannot Convert ID", result.Message)
		assert.False(t, result.Status)
	})
}

// MOCK SUCCESS
type mockCart struct {
}

//METHOD MOCK SUCCESS
func (m *mockCart) CreateCart(newAdd entities.Cart) (entities.Cart, error) {
	return entities.Cart{NameProduct: "Motor", NameSeller: "Otomotif Center"}, nil
}
func (m *mockCart) GetAllCart(UserID uint) ([]entities.Cart, []string, error) {
	return []entities.Cart{{NameProduct: "Motor", NameSeller: "Otomotif Center", ToBuy: "yes", Qty: 2, Price: 20000}, {NameProduct: "Mobil", NameSeller: "Otomotif Center"}}, []string{"Otomotif Center"}, nil
}
func (m *mockCart) GetCartID(x uint, UserID uint) (entities.Cart, error) {
	return entities.Cart{NameProduct: "Motor", NameSeller: "Otomotif Center"}, nil
}
func (m *mockCart) UpdateCart(id uint, updatedCart entities.Cart, UserID uint) (entities.Cart, error) {
	return entities.Cart{NameProduct: "Motor", NameSeller: "Otomotif Center"}, nil
}

func (m *mockCart) DeleteCart(id uint, UserID uint) error {
	return nil
}

func (m *mockCart) Shipment(UserID uint) (entities.Address, []entities.Cart, []string, error) {
	return entities.Address{}, []entities.Cart{}, []string{}, nil
}

// MOCK ERROR
type errMockCart struct {
}

// METHOD MOCK ERROR
func (e *errMockCart) CreateCart(newAdd entities.Cart) (entities.Cart, error) {
	return entities.Cart{}, errors.New("Access Database Error")
}

func (e *errMockCart) GetAllCart(UserID uint) ([]entities.Cart, []string, error) {
	return nil, []string{}, errors.New("Access Database Error")
}

func (e *errMockCart) GetCartID(x uint, UserID uint) (entities.Cart, error) {
	return entities.Cart{}, errors.New("Access Database Error")
}

func (e *errMockCart) UpdateCart(id uint, updatedCart entities.Cart, UserID uint) (entities.Cart, error) {
	return entities.Cart{}, errors.New("Access Database Error")
}

func (e *errMockCart) DeleteCart(id uint, UserID uint) error {
	return errors.New("Access Database Error")
}

func (e *errMockCart) Shipment(UserID uint) (entities.Address, []entities.Cart, []string, error) {
	return entities.Address{}, []entities.Cart{}, []string{}, errors.New("Access Database Error")
}
