package transaction

type InsertTransaction struct {
	Address       string `json:"address"`
	PaymentMethod string `json:"paymentMethod" validate:"required"`
}

type InsertStatusTransaction struct {
	Status string `json:"status" validate:"required"`
}
