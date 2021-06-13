package usecase

import "go-practice/pkg/domain"

type UserRepository interface {
	CreateUser(user domain.User) (err error)
	UpdateUser(user domain.User, id int) (err error)
	DeleteUser(id int) (err error)
	GetUser(id int) (user domain.User, err error)
	GetUserLists() ([]domain.User, error)
}
