package main

import (
	_ "github.com/lib/pq"
	"go-practice/pkg/infrastructure"
	"go-practice/pkg/interfaces/controllers"
	"go-practice/pkg/interfaces/database"
	"go-practice/pkg/usecase"
)

func main() {

	todo_repository := database.NewToDoRepository()
	todo_usecase := usecase.NewToDoUsecase(todo_repository)
	todo_controllers := controllers.NewToDoController(todo_usecase)

	user_repository := database.NewUserRepository()
	user_usecase := usecase.NewUserUsecase(user_repository)
	user_controllers := controllers.NewUserController(user_usecase)

	infrastructure.SetUpRouting(user_controllers, todo_controllers)
}
