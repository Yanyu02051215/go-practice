package usecase

import "go-practice/pkg/domain"

// interfaceに振る舞いを入れる
// interface型のUser
// ② userを作成している
type UserRepository interface {
	CreateUser(user domain.User) (err error)
	GetUser(id int) (user domain.User, err error)
}
