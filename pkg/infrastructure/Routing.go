package infrastructure

import (
	"github.com/gin-gonic/gin"
	"go-practice/pkg/interfaces/controllers"
)

func SetUpUserRouting(userController controllers.UsersController, todoController controllers.ToDosController) {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!!")
	})
	engine.GET("/users", userController.GetUsers())
	engine.POST("/users", userController.CreateUser())
	engine.GET("/user/:id", userController.GetUser())
	engine.PUT("/users/:id", userController.UpdateUser())
	engine.DELETE("/user/:id", userController.DeleteUser())
	engine.POST("/user/:id/todo", todoController.CreateToDo())
	engine.Run(":3000")
}
