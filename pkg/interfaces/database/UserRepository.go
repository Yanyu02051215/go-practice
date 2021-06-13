package database

import (
	"database/sql"
	"go-practice/pkg/domain"
	"go-practice/pkg/infrastructure"
	"go-practice/pkg/usecase"
	"log"

	_ "github.com/lib/pq"
)

type userRepository struct{}

func NewUserRepository() usecase.UserRepository {
	return userRepository{}
}

func (u userRepository) CreateUser(user domain.User) (err error) {
	infrastructure.DbCreate(`insert into users(id, name, email, password) values($1, $2, $3, $4);`, user.ID, user.Name, user.Email, user.Password)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u userRepository) UpdateUser(user domain.User, id int) (err error) {
	infrastructure.DbCreate(`update users set name = ?, email = ? where id = ?`, user.ID, user.Name, user.Email, user.Password)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u userRepository) DeleteUser(id int) (err error) {
	infrastructure.DbCreate(`delete from users where id = ?`, id)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u userRepository) GetUser(id int) (user domain.User, err error) {
	//infrastructure.DbCreate(`select id, name, email, password from users where id = ?`, user.ID, user.Name, user.Email, user.Password)
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=user password=Password! dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	cmd := `select id, name, email, password from users where id = ?`
	user = domain.User{}

	err = db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	return user, err
}

func (u userRepository) GetUserLists() ([]domain.User, error) {
	// infrastructure.DbSelect(`select id, name, email, password from users`)
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=user password=Password! dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	cmd := `select id, name, email, password from users`

	rows, err := db.Query(cmd)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, err
	// return []domain.User{}, nil
}
