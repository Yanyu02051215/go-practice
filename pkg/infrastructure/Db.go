package infrastructure

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func DbCreate(query string, args ...interface{}) (sql.Result, error) {
	//sqlパッケージのDBのポインタ型の変数
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=user password=Password! dbname=database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// cmd := `insert into users(id, name, email, password) values($1, $2, $3, $4);`

	// return db.Exec(query,
	// 	user.ID,
	// 	user.Name,
	// 	user.Email,
	// 	user.Password,
	// )
	return db.Exec(query, args)
}
