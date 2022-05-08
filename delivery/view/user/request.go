<<<<<<< HEAD
package user
=======
package user

type InsertUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UpdateUserRequest struct {
	Email string `json:"email" validate:"required"`
}
>>>>>>> 8fa9d85d2f209b23b12aedd593350d5089423d87
