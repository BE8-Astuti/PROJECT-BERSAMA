package transaction

// var token string

// // INITIATE TOKEN
// func TestCreateToken(t *testing.T) {
// 	t.Run("Create Token", func(t *testing.T) {
// 		token, _ = middlewares.CreateToken(1, "Motor", "Motor@gmail.com")
// 	})
// }

// func TestCreateTransaction(t *testing.T) {
// 	t.Run("Create Success", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"productId":   3,
// 			"nameSeller":  "Elec Center",
// 			"nameProduct": "Mouse",
// 			"qty":         3,
// 			"price":       100000,
// 			"toBuy":       "yes",
// 		})

// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart")
// 		TransactionC := NewRepoTrans(&mockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.CreateTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    interface{}
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 201, result.Code)
// 		assert.Equal(t, "Success Create Transaction", result.Message)
// 		assert.True(t, result.Status)
// 		assert.NotNil(t, result.Data)
// 	})
// 	t.Run("Error Access Database", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"productId":   3,
// 			"nameSeller":  "Elec Center",
// 			"nameProduct": "Mouse",
// 			"qty":         3,
// 			"price":       100000,
// 			"toBuy":       "yes",
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart")
// 		TransactionC := NewRepoTrans(&errMockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.CreateTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 500, result.Code)
// 		assert.Equal(t, "Cannot Access Database", result.Message)
// 		assert.False(t, result.Status)
// 	})
// 	t.Run("Error Bind", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody := "Jalan Gunung"
// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart")
// 		TransactionC := NewRepoTrans(&errMockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.CreateTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)
// 		assert.Equal(t, 415, result.Code)
// 		assert.Equal(t, "Cannot Bind Data", result.Message)
// 		assert.False(t, result.Status)
// 	})
// 	t.Run("Error Validate", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"productId":   3,
// 			"nameSeller":  "Elec Center",
// 			"nameProduct": "Mouse",
// 			"qty":         3,
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart")
// 		TransactionC := NewRepoTrans(&errMockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.CreateTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 406, result.Code)
// 		assert.Equal(t, "Validate Error", result.Message)
// 		assert.False(t, result.Status)
// 	})
// }

// // TEST GET ALL CART
// func TestGetAllTransaction(t *testing.T) {
// 	t.Run("Success Get All Transaction", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodPost, "/", nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart")
// 		GetTransaction := NewRepoTrans(&mockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.GetAllTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    interface{}
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 200, result.Code)
// 		assert.Equal(t, "Success Get All data", result.Message)
// 		assert.True(t, result.Status)
// 		assert.NotNil(t, result.Data)
// 	})
// 	t.Run("Error Access Database", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodPost, "/", nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart")
// 		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.GetAllTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 500, result.Code)
// 		assert.Equal(t, "Cannot Access Database", result.Message)
// 		assert.False(t, result.Status)
// 	})
// }

// // TEST UPDATE CART BY ID
// func TestUpdateTransaction(t *testing.T) {
// 	t.Run("Update Success", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"qty":   7,
// 			"toBuy": "no",
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart/:id")
// 		context.SetParamNames("id")
// 		context.SetParamValues("1")
// 		GetTransaction := NewRepoTrans(&mockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.UpdateTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    interface{}
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 200, result.Code)
// 		assert.Equal(t, "Updated", result.Message)
// 		assert.True(t, result.Status)
// 		assert.NotNil(t, result.Data)
// 	})
// 	t.Run("Error Not Found", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart/:id")
// 		context.SetParamNames("id")
// 		context.SetParamValues("7")
// 		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.UpdateTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 404, result.Code)
// 		assert.Equal(t, "Data Not Found", result.Message)
// 		assert.False(t, result.Status)
// 	})
// 	t.Run("Error Bind", func(t *testing.T) {
// 		e := echo.New()
// 		requestBody := "Jalan Gunung"
// 		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart/:id")
// 		context.SetParamNames("id")
// 		context.SetParamValues("7")
// 		TransactionC := NewRepoTrans(&errMockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(TransactionC.UpdateTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)
// 		assert.Equal(t, 415, result.Code)
// 		assert.Equal(t, "Cannot Bind Data", result.Message)
// 		assert.False(t, result.Status)
// 	})
// 	t.Run("Error Convert ID", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart/:id")
// 		context.SetParamNames("id")
// 		context.SetParamValues("C")
// 		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.UpdateTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 406, result.Code)
// 		assert.Equal(t, "Cannot Convert ID", result.Message)
// 		assert.False(t, result.Status)
// 	})
// }

// // TEST DELETE CART BY ID
// func TestDeleteTransaction(t *testing.T) {
// 	t.Run("Success Delete Transaction", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart/:id")
// 		context.SetParamNames("id")
// 		context.SetParamValues("7")
// 		GetTransaction := NewRepoTrans(&mockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.DeleteTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 200, result.Code)
// 		assert.Equal(t, "Deleted", result.Message)
// 		assert.True(t, result.Status)
// 	})
// 	t.Run("Error Delete Transaction", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart/:id")
// 		context.SetParamNames("id")
// 		context.SetParamValues("7")
// 		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.DeleteTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 500, result.Code)
// 		assert.Equal(t, "Cannot Access Database", result.Message)
// 		assert.False(t, result.Status)
// 	})
// 	t.Run("Error Convert ID", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/cart/:id")
// 		context.SetParamNames("id")
// 		context.SetParamValues("C")
// 		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New(), &MockSnap{})

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.DeleteTransaction())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 406, result.Code)
// 		assert.Equal(t, "Cannot Convert ID", result.Message)
// 		assert.False(t, result.Status)
// 	})
// }

// // TEST GET SHIPMENT DETAIL
// func TestShipment(t *testing.T) {
// 	t.Run("Success Get Shipment", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/shipment")

// 		GetTransaction := NewRepoTrans(&mockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.Shipment())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 			Data    interface{}
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 200, result.Code)
// 		assert.Equal(t, "Success Get Data Shipment", result.Message)
// 		assert.True(t, result.Status)
// 		assert.NotNil(t, result.Data)
// 	})
// 	t.Run("Error Get Shipment", func(t *testing.T) {
// 		e := echo.New()

// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/shipment")

// 		GetTransaction := NewRepoTrans(&errMockTransaction{}, validator.New())

// 		middleware.JWTWithConfig(middleware.JWTConfig{SigningMethod: "HS256", SigningKey: []byte("TOGETHER")})(GetTransaction.Shipment())(context)

// 		type Response struct {
// 			Code    int
// 			Message string
// 			Status  bool
// 		}

// 		var result Response
// 		json.Unmarshal([]byte(res.Body.Bytes()), &result)

// 		assert.Equal(t, 500, result.Code)
// 		assert.Equal(t, "Cannot Access Database", result.Message)
// 		assert.False(t, result.Status)
// 	})
// }

// type MockSnap struct {
// }

// // MOCK SUCCESS
// type mockTransaction struct {
// }

// //METHOD MOCK SUCCESS
// func (m *mockTransaction) CreateTransaction(NewTransaction entities.Transaction) (entities.Transaction, error) {
// 	return entities.Transaction{NameProduct: "Motor", NameSeller: "Otomotif Center"}, nil
// }
// func (m *mockTransaction) GetAllTransaction(UserID uint) ([]entities.Transaction, []string, error) {
// 	return []entities.Transaction{{NameProduct: "Motor", NameSeller: "Otomotif Center", ToBuy: "yes", Qty: 2, Price: 20000}, {NameProduct: "Mobil", NameSeller: "Otomotif Center"}}, []string{"Otomotif Center"}, nil
// }
// func (m *mockTransaction) GetTransactionDetail(UserID uint, OrderID string) (transaction.AllTrans, error) {
// 	return entities.Transaction{NameProduct: "Motor", NameSeller: "Otomotif Center"}, nil
// }
// func (m *mockTransaction) PayTransaction(UserID uint, OrderID string) (entities.Transaction, error) {
// 	return entities.Transaction{NameProduct: "Motor", NameSeller: "Otomotif Center"}, nil
// }

// func (m *mockTransaction) CancelTransaction(UserID uint, OrderID string) error {
// 	return nil
// }

// func (m *MockSnap) CreateTransaction(UserID uint) (entities.Address, []entities.Transaction, []string, error) {
// 	return entities.Address{Recipient: "Galih", HP: "123456"}, []entities.Transaction{{NameProduct: "Motor", NameSeller: "Otomotif Center", Qty: 2, Price: 50000, ToBuy: "yes"}}, []string{"Otomotif Center"}, nil
// }

// // MOCK ERROR
// type errMockTransaction struct {
// }

// // METHOD MOCK ERROR
// func (e *errMockTransaction) CreateTransaction(newAdd entities.Transaction) (entities.Transaction, error) {
// 	return entities.Transaction{}, errors.New("Access Database Error")
// }

// func (e *errMockTransaction) GetAllTransaction(UserID uint) ([]entities.Transaction, []string, error) {
// 	return nil, []string{}, errors.New("Access Database Error")
// }

// func (e *errMockTransaction) GetTransactionID(x uint, UserID uint) (entities.Transaction, error) {
// 	return entities.Transaction{}, errors.New("Access Database Error")
// }

// func (e *errMockTransaction) UpdateTransaction(id uint, updatedTransaction entities.Transaction, UserID uint) (entities.Transaction, error) {
// 	return entities.Transaction{}, errors.New("Access Database Error")
// }

// func (e *errMockTransaction) DeleteTransaction(id uint, UserID uint) error {
// 	return errors.New("Access Database Error")
// }

// func (e *errMockTransaction) Shipment(UserID uint) (entities.Address, []entities.Transaction, []string, error) {
// 	return entities.Address{}, []entities.Transaction{}, []string{}, errors.New("Access Database Error")
// }
