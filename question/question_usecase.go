package question

import (
	"creating-simple-api/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type QuestionUsecase interface {
	Insert(question model.Question) (*model.Question, error)
	FindAllQuestion(c *gin.Context) (*[]model.Question, error)
	FindByUUID(uuid uuid.UUID)(*model.Question, error)
	UpdateByUUID(uuid uuid.UUID, question model.Question) (*model.Question, error)
	DeleteByUUID(uuid uuid.UUID) error
}