package usecase

import (
	"creating-simple-api/constan"
	"creating-simple-api/middleware"
	"creating-simple-api/model"
	"creating-simple-api/user"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func CreateUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (e *UserUsecaseImpl) InserUser(user *model.User) (*model.User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		fmt.Printf("[UserUsecaseImpl.InsertUser] error bycrip password %v \n", err)
		return nil, fmt.Errorf(constan.FailedInsertData)
	}
	user.Password = string(pass)
	return e.userRepo.InserUser(user)
}

func (e *UserUsecaseImpl) FindUserByEmail(email string) (*model.User, error) {
	return e.userRepo.FindUserByEmail(email)
}

func (e *UserUsecaseImpl) FindUserById(id int) (*model.User, error) {
	return e.userRepo.FindUserById(id)
}

func (e *UserUsecaseImpl) ValidUser(c *gin.Context) (*model.User, error) {
	authDetail, err := middleware.ExtractTokenAuth(c)
	if err != nil {
		return nil, err
	}
	user, err := e.userRepo.FindUserById(authDetail.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}


