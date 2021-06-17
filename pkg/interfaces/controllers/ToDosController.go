package controllers

import (
	"go-practice/pkg/domain"
	"go-practice/pkg/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDosController interface {
	CreateToDo() gin.HandlerFunc
}

type toDosController struct {
	usecase usecase.ToDoUseCase
}

func NewToDoController(usecase usecase.ToDoUseCase) ToDosController {
	return &toDosController{
		usecase: usecase,
	}
}

func (t toDosController) CreateToDo() gin.HandlerFunc {
	return func(c *gin.Context) {
		todo := domain.Todo{
			ID:      1,
			Title:   "テストタイトル",
			Context: "テストコンテキスト",
			UserId:  2,
		}

		err := t.usecase.RegisterTodo(todo)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": todo,
		})
	}
}
