package database

import (
	// "database/sql"
	"go-practice/pkg/domain"
	"go-practice/pkg/infrastructure"
	"go-practice/pkg/usecase"
	"log"

	_ "github.com/lib/pq"
)

type todoRepository struct{}

func NewToDoRepository() usecase.ToDoRepository {
	return todoRepository{}
}

func (u todoRepository) CreateToDo(todo domain.Todo) (err error) {
	infrastructure.DbCreate(`insert into todos(id, title, context, user_id) values($1, $2, $3, $4);`, todo.ID, todo.Title, todo.Context, todo.UserId)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
