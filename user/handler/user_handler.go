package handler

import (
	"creating-simple-api/constan"
	"creating-simple-api/model"
	"creating-simple-api/user"
	"creating-simple-api/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type UserHandler struct {
	userUsecase user.UserUsecase
}

func CreateUserHandler(r *gin.Engine, userUsecase user.UserUsecase) {
	userHandler := UserHandler{userUsecase}

	r.POST("/user/register", userHandler.register)
	r.POST("/user/login", userHandler.login)
}

func (e *UserHandler) register(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Printf("[UserHandler.register] error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, constan.FailedInternalServerError)
		return
	}
	if user.Email == "" || user.Name == "" || user.Password == "" {
		utils.HandleError(c, http.StatusBadRequest, constan.FieldsAreRequired)
		return
	}
	_, err = e.userUsecase.FindUserByEmail(user.Email)
	if err == nil {
		utils.HandleError(c, http.StatusBadRequest, constan.EmailIsExsis)
		return
	}
	output, err := e.userUsecase.InserUser(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, output)
}

func (e *UserHandler) login(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Printf("[UserHandler.login] error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, constan.FailedInternalServerError)
		return
	}
	mUser, err := e.userUsecase.FindUserByEmail(user.Email)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(mUser.Password), []byte(user.Password))
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, constan.PasswordWrong)
		return
	}
	token, err := utils.GenerateToken(int(mUser.ID), mUser.Password)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
	}
	var output = model.Auth{
		Token: token,
	}
	utils.HandleSuccess(c, output)
}