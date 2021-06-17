package usecase

import "go-practice/pkg/domain"

type ToDoRepository interface {
	CreateToDo(todo domain.Todo) (err error)
}
