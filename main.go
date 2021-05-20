package main

import (
	"database/sql"
	"go-practice/pkg/infras"
	"go-practice/pkg/interfaces"
	"go-practice/pkg/usecases"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	Db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=user password=Password! dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer Db.Close()

	// func GetUser(id int) (user User, err error) {
	// 	user = User{}
	// 	cmd := `select id, uuid, name, email, password
	// 	from users where id = ?`
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

	repository := infras.NewUser()
	usecase := usecases.NewUser(repository)
	interfaces := interfaces.NewUser(usecase)

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!!")
	})
	// engine.GET("/users", interfaces.GetUsers())
	engine.POST("/users", interfaces.CreateUser())
	engine.Run(":3000")

}
