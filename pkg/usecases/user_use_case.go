package usecases

import (
	"go-practice/pkg/domains/models"
	"go-practice/pkg/domains/repository"
)

type UserUseCase interface {
	CreateUser(user models.User) (err error)
}

// interface型のUserを入れてる
// repository.UserはinfraのCreateUserを指示したinterface
type userUseCase struct {
	repository repository.UserRepository
}

// ④ 上にあるUserのinterfaceを返り値にしている(それにより下のCreateUserが発火している)
// repository.UserはinfraのCreateUserを指示したinterface
// 引数のrepositoryはrepositoryのUserで作られたからrepository.Userの構造体になっているはず??
func NewUser(repository repository.UserRepository) UserUseCase {
	return &userUseCase{
		repository: repository,
	}
}

// ここの引数や返り値は基本的に上記のintefaceと同じ
func (u userUseCase) CreateUser(user models.User) (err error) {
	// このCreateUserはrepositoryにあるCreateUser
	err = u.repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

// func (u user) GetUsers(id models.User) (user repository.User, err error) {
// 	err = u.repository.GetUsers(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
