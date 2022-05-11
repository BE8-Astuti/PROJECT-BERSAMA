package transaction

type InsertTransaction struct {
	Address string `json:"address" validate:"required"`
}

type InsertStatusTransaction struct {
	Status string `json:"status" validate:"required"`
}
