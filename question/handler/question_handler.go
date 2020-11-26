package handler

import (
	"creating-simple-api/constan"
	"creating-simple-api/middleware"
	"creating-simple-api/model"
	"creating-simple-api/question"
	"creating-simple-api/user"
	"creating-simple-api/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type QuestionHandler struct {
	questionUsecase question.QuestionUsecase
	userUsecase user.UserUsecase
}

func CreateQuestionHandler(r *gin.Engine, questionUsecase question.QuestionUsecase, userUsecase user.UserUsecase) {
	questionHandler := QuestionHandler{questionUsecase, userUsecase}

	r2 := r.Group("/question").Use(middleware.TokenVerifikasiMiddleware())
	r2.POST("", questionHandler.addQuestion)
	r2.GET("", questionHandler.FindAllQuestion)
	r2.GET("/:uuid", questionHandler.FindByUUID)
	r2.PUT("/:uuid", questionHandler.UpdateQuestion)
	r2.DELETE("/:uuid", questionHandler.deleteQuestion)
}

func (e *QuestionHandler) addQuestion(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	var question model.Question
	err = c.Bind(&question)
	if err != nil {
		fmt.Printf("[QuestionHandler.addQuestion] error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, constan.FailedInternalServerError)
		return
	}
	if question.Question == "" {
		utils.HandleError(c, http.StatusBadRequest, constan.FieldsAreRequired)
		return
	}
	question.CreatedBy = user.Name
	question.UpdateBy = user.Name
	question.IsActive = true
	output, err := e.questionUsecase.Insert(question)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, output)
}

func (e *QuestionHandler) FindAllQuestion(c *gin.Context) {
	question, err := e.questionUsecase.FindAllQuestion(c)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, question)
}

func (e *QuestionHandler) FindByUUID(c *gin.Context) {
	uid := c.Param("uuid")
	uuid, err := uuid.Parse(uid)
	if err != nil {
		fmt.Printf("[QuestionHandler.FindByUUID] error convert string to uuid %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, constan.FailedInternalServerError)
		return
	}
	question, err := e.questionUsecase.FindByUUID(uuid)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSuccess(c, question)
}

func (e *QuestionHandler) UpdateQuestion(c *gin.Context) {
	user, err := e.userUsecase.ValidUser(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, err.Error())
		return
	}
	uid := c.Param("uuid")
	uuid, err := uuid.Parse(uid)
	if err != nil {
		fmt.Printf("[QuestionHandler.UpdateQuestion] error conver string to uuid %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, constan.FailedInternalServerError)
		return
	}
	_, err = e.questionUsecase.FindByUUID(uuid)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	} 
	var question model.Question
	err = c.Bind(&question)
	if err != nil {
		fmt.Printf("[QuestionHandler.updateQuestion] error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, constan.FailedInternalServerError)
		return
	}
	if question.Question == "" {
		utils.HandleError(c, http.StatusBadRequest, constan.FieldsAreRequired)
		return
	}
	question.UpdateBy = user.Name
	output, err := e.questionUsecase.UpdateByUUID(uuid, question)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c,output)
}

func (e *QuestionHandler) deleteQuestion(c *gin.Context) {
	uid := c.Param("uuid")
	uuid, err := uuid.Parse(uid)
	if err != nil {
		fmt.Printf("[QuestionHandler.deleteQuestion] error conver string to uuid %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, constan.FailedInternalServerError)
		return
	}
	_, err = e.questionUsecase.FindByUUID(uuid)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	err = e.questionUsecase.DeleteByUUID(uuid)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, constan.SucessDeleteData)
}