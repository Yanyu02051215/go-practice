package infrastructure

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=user password=Password! dbname=database sslmode=disable")
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
