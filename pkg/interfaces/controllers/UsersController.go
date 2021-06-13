package controllers

import (
	"go-practice/pkg/domain"
	"go-practice/pkg/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController interface {
	CreateUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
}

type usersController struct {
	usecase usecase.UserUseCase
}

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

func (u usersController) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := domain.User{
			ID:       2,
			Name:     "JIRO",
			Email:    "changed@gmail.com",
			Password: "password2",
		}
		id := c.Params.ByName("id")
		idInt, _ := strconv.Atoi(id)

		err := u.usecase.ReRegisterUser(user, idInt)
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

func (u usersController) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		idInt, _ := strconv.Atoi(id)

		err := u.usecase.DeleteUser(idInt)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "deleted",
		})
	}
}

func (u usersController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		idInt, _ := strconv.Atoi(id)
		user, err := u.usecase.GetUser(idInt)

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

func (u usersController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		users, err := u.usecase.GetUsers()
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": users,
		})
	}
}
