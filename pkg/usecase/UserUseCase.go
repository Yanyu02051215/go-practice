package usecase

import (
	"go-practice/pkg/domain"
)

type UserUseCase interface {
	RegisterUser(user domain.User) (err error)
	ReRegisterUser(user domain.User, id int) (err error)
	DeleteUser(id int) (err error)
	GetUser(id int) (user domain.User, err error)
	GetUsers() ([]domain.User, error)
}

type userUseCase struct {
	repository UserRepository
}

// UserUseCaseはUserRepositoryに依存してる
func NewUserUsecase(repository UserRepository) UserUseCase {
	return &userUseCase{
		repository: repository,
	}
}

func (u userUseCase) RegisterUser(user domain.User) (err error) {
	return u.repository.CreateUser(user)
}

func (u userUseCase) ReRegisterUser(user domain.User, id int) (err error) {
	return u.repository.UpdateUser(user, id)
}

func (u userUseCase) DeleteUser(id int) (err error) {
	return u.repository.DeleteUser(id)
}

func (u userUseCase) GetUser(id int) (user domain.User, err error) {
	return u.repository.GetUser(id)
}

func (u userUseCase) GetUsers() ([]domain.User, error) {
	return u.repository.GetUserLists()
}
