package product

type InsertProdukRequest struct {
	UserID      uint   `json:"user_id" validate:"required"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name" validate:"required"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type UpdateProdukRequest struct {
	Name        string `json:"name" validate:"required"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
