package infras

import (
	"database/sql"
	"go-practice/pkg/domains/models"
	"go-practice/pkg/domains/repository"
	"log"

	_ "github.com/lib/pq"
)

type user struct{}

var Db *sql.DB
var er error

func init() {
	Db, er = sql.Open("postgres", "host=127.0.0.1 port=5555 user=user password=Password! dbname=database sslmode=disable")
	if er != nil {
		log.Fatal(er)
	}

	er = Db.Ping()
	if er != nil {
		log.Fatal(er)
	}
}

// ①repositoryのUser(interface)を返り値にしている  又、返り値をuser構造体の中に入れていってる
func NewUser() repository.User {
	return &user{}
}

// ③repositoryのinterfaceからの振る舞いでメソッドを発火させてる
func (u user) CreateUser(user models.User) (err error) {
	// Db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=user password=Password! dbname=database sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = Db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	defer Db.Close()

	cmd := `insert into users(id, name, email, password) values($1, $2, $3, $4);`

	_, err = Db.Exec(cmd,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
