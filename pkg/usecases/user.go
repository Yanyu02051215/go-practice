package usecases

import (
	"go-practice/pkg/domains/models"
	"go-practice/pkg/domains/repository"
)

type User interface {
	CreateUser(user models.User) (err error)
}

// interface型のUserを入れてる
// repository.UserはinfraのCreateUserを指示したinterface
type user struct {
	repository repository.User
}

// ④ 上にあるUserのinterfaceを返り値にしている(それにより下のCreateUserが発火している)
// repository.UserはinfraのCreateUserを指示したinterface
// 引数のrepositoryはrepositoryのUserで作られたからrepository.Userの構造体になっているはず??
func NewUser(repository repository.User) User {
	return &user{
		repository: repository,
	}
}

// ここの引数や返り値は基本的に上記のintefaceと同じ
func (u user) CreateUser(user models.User) (err error) {
	// このCreateUserはrepositoryにあるCreateUser
	err = u.repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
