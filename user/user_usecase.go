package user

import (
	"creating-simple-api/model"
	"github.com/gin-gonic/gin"
)


type UserUsecase interface {
	InserUser(user *model.User) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserById(id int) (*model.User, error)
	ValidUser(c *gin.Context) (*model.User, error)
}