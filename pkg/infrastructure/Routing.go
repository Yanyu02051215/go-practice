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
	// engine.GET("/users", interfaces.GetUsers())
	engine.GET("/users", userController.CreateUser())
	engine.Run(":3000")
}
