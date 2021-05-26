package repository

import "go-practice/pkg/domains/models"

// interfaceに振る舞いを入れる
// interface型のUser
// ② userを作成している
type UserRepository interface {
	CreateUser(user models.User) (err error)
	GetUsers(id models.User) (user models.User, err error)
}
