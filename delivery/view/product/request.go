package product

type InsertProdukRequest struct {
	CategoryID  int    `json:"category_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateProdukRequest struct {
	Name        string `json:"name" validate:"required"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
