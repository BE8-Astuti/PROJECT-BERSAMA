package category

type InsertCat struct {
	UserID int    `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type UpdateCat struct {
	Name string `json:"name"`
}
