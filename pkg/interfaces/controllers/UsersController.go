package controllers

import (
	"go-practice/pkg/domain"
	"go-practice/pkg/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController interface {
	CreateUser() gin.HandlerFunc
}

type usersController struct {
	usecase usecase.UserUseCase
}

// 引数のusecasesはusecaseのCreateUserで作られたからusecases.Userの構造体になっているはず??
func NewUserController(usecase usecase.UserUseCase) UsersController {
	return &usersController{
		usecase: usecase,
	}
}

func (u usersController) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := domain.User{
			ID:       1,
			Name:     "JIRO",
			Email:    "examole2@gmail.com",
			Password: "password2",
		}

		err := u.usecase.RegisterUser(user)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": user,
		})
	}
}
