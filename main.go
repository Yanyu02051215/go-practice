package main

import (
	_ "github.com/lib/pq"
	"go-practice/pkg/infrastructure"
	"go-practice/pkg/interfaces/controllers"
	"go-practice/pkg/interfaces/database"
	"go-practice/pkg/usecase"
)

func main() {

	repository2 := database.NewToDoRepository()
	usecase2 := usecase.NewToDoUsecase(repository2)
	controllers2 := controllers.NewToDoController(usecase2)

	repository := database.NewUserRepository()
	usecase := usecase.NewUserUsecase(repository)
	controllers := controllers.NewUserController(usecase)

	infrastructure.SetUpRouting(controllers, controllers2)
}
