package usecase

import (
	"go-practice/pkg/domain"
)

type ToDoUseCase interface {
	RegisterTodo(todo domain.Todo) (err error)
}

type todoUseCase struct {
	repository ToDoRepository
}

// UserUseCaseはUserRepositoryに依存してる
func NewToDoUsecase(repository ToDoRepository) ToDoUseCase {
	return &todoUseCase{
		repository: repository,
	}
}

func (t todoUseCase) RegisterTodo(todo domain.Todo) (err error) {
	return t.repository.CreateToDo(todo)
}
