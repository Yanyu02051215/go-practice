package infras

import (
	// "database/sql"
	"go-practice/pkg/domains/models"
	"go-practice/pkg/domains/repository"
	"log"

	_ "github.com/lib/pq"
)

// domein領域が持っていないといけないものがdomainディレクトリ以外にあるのは良くない(domainの知識ではなくなる)
//userRepositoryに定義されているメソッドを持つ

// メソッドを持たせるために定義している
type userRepositoryImpl struct{}

// ①repositoryのUser(interface)を返り値にしている  又、返り値をuser構造体の中に入れていってる
func NewUser() repository.UserRepository {
	return userRepositoryImpl{}
}

// ③repositoryのinterfaceからの振る舞いでメソッドを発火させてる
func (u userRepositoryImpl) CreateUser(user models.User) (err error) {
	DbCreate(`insert into users(id, name, email, password) values($1, $2, $3, $4);`, user.ID, user.Name, user.Email, user.Password)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u userRepositoryImpl) GetUsers(id models.User) (user models.User, err error) {
	DbCreate(`select id, name, email, password from users where id = ?`, user.ID, user.Name, user.Email, user.Password)
	// cmd := `select id, name, email, password from users where id = ?`
	// err = db.QueryRow(cmd, id).Scan(
	// 	&user.ID,
	// 	&user.Name,
	// 	&user.Email,
	// 	&user.Password,
	// )
	return user, err
}
