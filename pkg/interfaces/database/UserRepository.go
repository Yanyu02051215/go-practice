package database

import (
	// "database/sql"
	"go-practice/pkg/domain"
	"go-practice/pkg/infrastructure"
	"go-practice/pkg/usecase"
	"log"

	_ "github.com/lib/pq"
)

// domain領域が持っていないといけないものがdomainディレクトリ以外にあるのは良くない(domainの知識ではなくなる)
//userRepositoryに定義されているメソッドを持つ

// メソッドを持たせるために定義している
type userRepository struct{}

// ①repositoryのUser(interface)を返り値にしている  又、返り値をuser構造体の中に入れていってる
func NewUserRepository() usecase.UserRepository {
	return userRepository{}
}

// ③repositoryのinterfaceからの振る舞いでメソッドを発火させてる
func (u userRepository) CreateUser(user domain.User) (err error) {
	infrastructure.DbCreate(`insert into users(id, name, email, password) values($1, $2, $3, $4);`, user.ID, user.Name, user.Email, user.Password)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u userRepository) GetUser(id int) (user domain.User, err error) {
	infrastructure.DbCreate(`select id, name, email, password from users where id = ?`, user.ID, user.Name, user.Email, user.Password)
	// cmd := `select id, name, email, password from users where id = ?`
	// err = db.QueryRow(cmd, id).Scan(
	// 	&user.ID,
	// 	&user.Name,
	// 	&user.Email,
	// 	&user.Password,
	// )
	return user, err
}
