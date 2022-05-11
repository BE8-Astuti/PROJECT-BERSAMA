package user

import "together/be8/entities"

type User interface {
	InsertUser(newUser entities.User) (entities.User, error)
<<<<<<< HEAD
	// GetAllUser() ([]entities.User, error)
=======
>>>>>>> fbedbae8ed32763c12abe5f92c0cbd8da656f0dc
	GetUserID(ID int) (entities.User, error)
	UpdateUser(ID int, email string) (entities.User, error)
	DeleteUser(ID int) (entities.User, error)
	Login(email string, password string) (entities.User, error)
}
