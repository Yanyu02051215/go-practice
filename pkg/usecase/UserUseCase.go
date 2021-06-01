package usecase

import (
	"go-practice/pkg/domain"
)

type UserUseCase interface {
	RegisterUser(user domain.User) (err error)
}

// interface型のUserを入れてる
// repository.UserはinfraのCreateUserを指示したinterface
type userUseCase struct {
	repository UserRepository
}

// ④ 上にあるUserのinterfaceを返り値にしている(それにより下のCreateUserが発火している)
// repository.UserはinfraのCreateUserを指示したinterface
// 引数のrepositoryはrepositoryのUserで作られたからrepository.Userの構造体になっているはず??
func NewUserUsecase(repository UserRepository) UserUseCase {
	return &userUseCase{
		repository: repository,
	}
}

// ここの引数や返り値は基本的に上記のintefaceと同じ
func (u userUseCase) RegisterUser(user domain.User) (err error) {
	// このCreateUserはrepositoryにあるCreateUser
	return u.repository.CreateUser(user)
}

func (u userUseCase) GetUser(id int) (user domain.User, err error) {
	return u.repository.GetUser(id)
}
