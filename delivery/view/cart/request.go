package cart

type InsertCart struct {
	NameSeller  string `json:"nameSeller" validate:"required"`
	NameProduct string `json:"nameProduct" validate:"required"`
	Qty         int    `json:"qty" validate:"required"`
	Price       int    `json:"price" validate:"required"`
<<<<<<< HEAD
	ToBuy       string `json:"toBuy" validate:"required"`
	ProductID   uint   `json:"productId" validate:"required"`
}

type CartProduct struct {
	NameProduct string `json:"nameProduct" validate:"required"`
	Qty         int    `json:"qty" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	ToBuy       string `json:"toBuy" validate:"required"`
}

type GetCart struct {
	NameSeller string        `json:"nameSeller"`
	Product    []CartProduct `json:"product"`
}

=======
	ProductID   uint   `json:"productId" validate:"required"`
}

>>>>>>> fbedbae8ed32763c12abe5f92c0cbd8da656f0dc
type UpdateCart struct {
	Qty   int    `json:"qty"`
	ToBuy string `json:"toBuy"`
}
