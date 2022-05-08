package cart

type InsertCart struct {
	NameSeller  string `json:"nameSeller" validate:"required"`
	NameProduct string `json:"nameProduct" validate:"required"`
	Qty         int    `json:"qty" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	ToBuy       bool   `json:"toBuy" validate:"required"`
}

type CartProduct struct {
	NameProduct string `json:"nameProduct" validate:"required"`
	Qty         int    `json:"qty" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	ToBuy       bool   `json:"toBuy" validate:"required"`
}

type GetCart struct {
	NameSeller string        `json:"nameSeller"`
	Product    []CartProduct `json:"product"`
}
