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

	return db.Exec(query, args)
}

// func DbQuery(query string) (rows, error) {
// 	//sqlパッケージのDBのポインタ型の変数
// 	db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=user password=Password! dbname=database sslmode=disable")

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()

// 	rows, err := db.Query(query)

// 	return rows, err
// }
