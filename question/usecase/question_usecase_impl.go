package usecase

import (
	"creating-simple-api/model"
	"creating-simple-api/question"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type QuestionUsecaseImpl struct {
	questionRepo question.QuestionRepo
}

func CreateQuestionUsecase(questionRepo question.QuestionRepo) question.QuestionUsecase {
	return &QuestionUsecaseImpl{questionRepo}
}

func (e *QuestionUsecaseImpl) Insert(question model.Question) (*model.Question, error) {
	uid, err :=  uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	question.UUID = uid
	return e.questionRepo.Insert(question)
}

func (e *QuestionUsecaseImpl) FindAllQuestion(c *gin.Context) (*[]model.Question, error) {
	return e.questionRepo.FindAllQuestion(c)
}

func (e *QuestionUsecaseImpl) FindByUUID(uuid uuid.UUID)(*model.Question, error) {
	return e.questionRepo.FindByUUID(uuid)
}

func (e *QuestionUsecaseImpl) UpdateByUUID(uuid uuid.UUID, question model.Question) (*model.Question, error) {
	return e.questionRepo.UpdateByUUID(uuid, question)
}

func (e *QuestionUsecaseImpl) DeleteByUUID(uuid uuid.UUID) error {
	return e.questionRepo.DeleteByUUID(uuid)
}