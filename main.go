package main

import (
	_ "github.com/lib/pq"
	"go-practice/pkg/infrastructure"
	"go-practice/pkg/interfaces/controllers"
	"go-practice/pkg/interfaces/database"
	"go-practice/pkg/usecase"
)

func main() {

	// Db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=user password=Password! dbname=database sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = Db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer Db.Close()

	// func GetUser(id int) (user User, err error) {
	// 	user = User{}
	// 	cmd := `select id, uuid, name, email, password from users where id = ?`
	// 	err = Db.QueryRow(cmd, id).Scan(
	// 		&user.ID,
	// 		&user.Name,
	// 		&user.Email,
	// 		&user.Password,
	// 	)
	// 	return user, err
	// }

	// u, err := GetUser(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	repository := database.NewUserRepository()
	usecase := usecase.NewUserUsecase(repository)
	controllers := controllers.NewUserController(usecase)

	infrastructure.SetUpRouting(controllers)
}
