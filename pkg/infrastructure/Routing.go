package infrastructure

import (
	"github.com/gin-gonic/gin"
	"go-practice/pkg/interfaces/controllers"
)

func SetUpRouting(userController controllers.UsersController) {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin!!")
	})
	engine.GET("/users", userController.GetUsers())
	engine.GET("/user/:id", userController.GetUser())
	engine.PUT("/users/:id", userController.UpdateUser())
	engine.DELETE("/user/:id", userController.DeleteUser())
	engine.Run(":3000")
}
