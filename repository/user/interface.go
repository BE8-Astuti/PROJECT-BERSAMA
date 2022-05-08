package user

import "together/be8/entities"

type User interface {
	InsertUser(newUser entities.User) (entities.User, error)
	GetAllUser() ([]entities.User, error)
	GetUserID(ID int) (entities.User, error)
	UpdateUser(ID int, email string) (entities.User, error)
	DeleteUser(ID int) (entities.User, error)
	Login(email string, password string) (entities.User, error)
}
