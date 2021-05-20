package interfaces

import (
	"go-practice/pkg/domains/models"
	"go-practice/pkg/usecases"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User interface {
	CreateUser() gin.HandlerFunc
}

type user struct {
	usecase usecases.User
}

// 引数のusecasesはusecaseのCreateUserで作られたからusecases.Userの構造体になっているはず??
func NewUser(usecase usecases.User) User {
	return &user{
		usecase: usecase,
	}
}

func (u user) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{
			ID:       1,
			Name:     "JIRO",
			Email:    "examole2@gmail.com",
			Password: "password2",
		}
		// 処理って毎回行われてるの??
		err := u.usecase.CreateUser(user)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "careted",
		})
	}
}
