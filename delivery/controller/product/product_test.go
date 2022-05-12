package product

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
		token, _ = middlewares.CreateToken(1, "yani", "yani@gmail.com")
	})
}

func TestInsertProd(t *testing.T) {
	t.Run("Create Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"categoryId":  1,
			"nameSeller":  "Xiaomi",
			"name":        "Poco X3",
			"stock":       17,
			"price":       1500000,
			"description": "Barang Bagus",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product")
		Product := New(&mockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Product.InsertProd())(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response
		json.Unmarshal([]byte(res.Body.Bytes()), &result)

		assert.Equal(t, 201, result.Code)
		assert.Equal(t, "Success Create Product", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Access Database", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"categoryId":  1,
			"nameSeller":  "Xiaomi",
			"name":        "Poco X3",
			"stock":       17,
			"price":       1500000,
			"description": "Barang Bagus",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product")
		Prod := New(&errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Prod.InsertProd())(context)

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

		requestBody := "kecantikan"

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product")
		Prod := New(&errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Prod.InsertProd())(context)

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

			"category_id": 1,
			"name":        "tango",
			"stock":       10,
			"price":       10000,
			"description": "wafer coklat",
		})

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product")
		Prod := New(&errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Prod.InsertProd())(context)

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

func TestGetAllProduct(t *testing.T) {
	t.Run("Success Get All Product", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product")

		Produk := New(&mockProduct{}, validator.New())

		Produk.GetAllProduct()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response

		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All Data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Success Get All Product", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product")

		Produk := New(&errMockProduct{}, validator.New())

		Produk.GetAllProduct()(context)

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

func TestGetProdbyID(t *testing.T) {
	t.Run("Success Get Produk By ID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		Produk := New(&mockProduct{}, validator.New())

		Produk.GetProID()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response

		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get Data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		Produk := New(&errMockProduct{}, validator.New())

		Produk.GetProID()(context)

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

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		produk := New(&errMockProduct{}, validator.New())

		produk.GetProID()(context)

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

func TestGetProdukbySeller(t *testing.T) {
	t.Run("Success Get Produk By Seller", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		Produk := New(&mockProduct{}, validator.New())

		Produk.GetProdukbySeller()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response

		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get Data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		Produk := New(&errMockProduct{}, validator.New())

		Produk.GetProdukbySeller()(context)

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

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/user/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		produk := New(&errMockProduct{}, validator.New())

		produk.GetProdukbySeller()(context)

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

func TestGetProdukByCategory(t *testing.T) {
	t.Run("Success Get Produk By Category", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		Produk := New(&mockProduct{}, validator.New())

		Produk.GetProdukByCategory()(context)

		type Response struct {
			Code    int
			Message string
			Status  bool
			Data    interface{}
		}

		var result Response

		json.Unmarshal([]byte(res.Body.Bytes()), &result)
		assert.Equal(t, 200, result.Code)
		assert.Equal(t, "Success Get All Data", result.Message)
		assert.True(t, result.Status)
		assert.NotNil(t, result.Data)
	})
	t.Run("Error Not Found", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		Produk := New(&errMockProduct{}, validator.New())

		Produk.GetProdukByCategory()(context)

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

		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		produk := New(&errMockProduct{}, validator.New())

		produk.GetProdukByCategory()(context)

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

func TestUpdateProduk(t *testing.T) {
	t.Run("Update Success", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":        "tango coklat",
			"stock":       10,
			"price":       10000,
			"description": "wafer coklat enak",
		})
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		Prod := New(&mockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Prod.UpdateProduk())(context)

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
		context.SetPath("/product/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		Prod := New(&errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Prod.UpdateProduk())(context)

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
		requestBody := "Tango"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		Prod := New(&errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Prod.UpdateProduk())(context)

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
		context.SetPath("/address/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		Prod := New(&errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Prod.UpdateProduk())(context)

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

func TestDeleteProduk(t *testing.T) {
	t.Run("Success Delete Produk", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		produk := New(&mockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(produk.DeleteProduk())(context)

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
	t.Run("Error Delete Produk", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/product/:id")
		context.SetParamNames("id")
		context.SetParamValues("7")
		Prod := New(&errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Prod.DeleteProduk())(context)

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
		context.SetPath("/product/:id")
		context.SetParamNames("id")
		context.SetParamValues("C")
		Prod := New(&errMockProduct{}, validator.New())

		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(Prod.DeleteProduk())(context)

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
type mockProduct struct {
}

//METHOD MOCK SUCCESS
func (s *mockProduct) InsertProduk(newProduk entities.Product) (entities.Product, error) {
	return entities.Product{UserID: 1, CategoryID: 1, Name: "tango", Stock: 10, Price: 10000, Description: "wafer coklat"}, nil
}

func (s *mockProduct) GetAllProduct() ([]entities.Product, error) {
	return []entities.Product{{Name: "Baju"}, {Name: "Celana"}}, nil
}

func (s *mockProduct) GetProdbyID(id uint) (entities.Product, error) {
	return entities.Product{UserID: 1, CategoryID: 1, Name: "tango", Stock: 10, Price: 10000, Description: "wafer coklat"}, nil
}
func (s *mockProduct) GetProdBySeller(UserID uint) ([]entities.Product, error) {
	return []entities.Product{{UserID: 1, CategoryID: 1, Name: "tango", Stock: 10, Price: 10000, Description: "wafer coklat"}, {UserID: 1, CategoryID: 1, Name: "roti", Stock: 10, Price: 10000, Description: "roti kasur"}}, nil
}
func (s *mockProduct) GetProdByCategory(id int) ([]entities.Product, error) {
	return []entities.Product{{UserID: 1, CategoryID: 1, Name: "roti", Stock: 10, Price: 10000, Description: "roti kasur"}}, nil
}

func (s *mockProduct) UpdateProduk(id int, UpdateProduk entities.Product, UserID uint) (entities.Product, error) {
	return entities.Product{UserID: 1, CategoryID: 1, Name: "tango", Stock: 10, Price: 10000, Description: "wafer coklat"}, nil
}
func (s *mockProduct) DeleteProduk(id uint, UserID uint) error {
	return nil
}

// MOCK ERROR
type errMockProduct struct{}

func (d *errMockProduct) InsertProduk(newProduk entities.Product) (entities.Product, error) {
	return entities.Product{}, errors.New("Access Database Error")
}

func (d *errMockProduct) GetAllProduct() ([]entities.Product, error) {
	return []entities.Product{}, errors.New("Access Database Error")
}

func (d *errMockProduct) GetProdbyID(id uint) (entities.Product, error) {
	return entities.Product{}, errors.New("Access Database Error")
}

func (d *errMockProduct) GetProdBySeller(UserID uint) ([]entities.Product, error) {
	return nil, errors.New("Access Database Error")
}

func (d *errMockProduct) GetProdByCategory(id int) ([]entities.Product, error) {
	return nil, errors.New("Access Database Error")
}

func (d *errMockProduct) UpdateProduk(id int, UpdateProduk entities.Product, UserID uint) (entities.Product, error) {
	return entities.Product{}, errors.New("Access Database Error")
}

func (d *errMockProduct) DeleteProduk(id uint, UserID uint) error {
	return errors.New("Access Database Error")
}
