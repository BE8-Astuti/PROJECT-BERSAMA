package cart

type InsertCart struct {
	ProductID uint `json:"productId" validate:"required"`
	Qty       int  `json:"qty" validate:"required"`
}

type UpdateCart struct {
	Qty   int    `json:"qty"`
	ToBuy string `json:"toBuy"`
}
