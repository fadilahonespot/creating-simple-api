package repo

import (
	"creating-simple-api/constan"
	"creating-simple-api/model"
	"creating-simple-api/question"
	"creating-simple-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	"golang.org/x/exp/errors/fmt"
)


type QuestionRepoImpl struct {
	db *gorm.DB
}

func CreateQuestionRepo(db *gorm.DB) question.QuestionRepo {
	return &QuestionRepoImpl{db}
}

func (e *QuestionRepoImpl) Insert(question model.Question) (*model.Question, error) {
	err := e.db.Save(&question).Error
	if err != nil {
		fmt.Printf("[QuestionRepoImpl.Insert]error execute query %v \n", err)
		return nil, fmt.Errorf(constan.FailedInsertData)
	}
	return &question, nil
}

func (e *QuestionRepoImpl) FindAllQuestion(c *gin.Context) (*[]model.Question, error) {
	var question []model.Question
	err := e.db.Scopes(utils.Paginate(c)).Find(&question).Error
	if err != nil {
		fmt.Printf("[QuestionRepoImpl.FindAllQuestion] error execute query %v \n", err)
		return nil, err
	}
	return &question, nil
}

func (e *QuestionRepoImpl) FindByUUID(uuid uuid.UUID)(*model.Question, error) {
	var question model.Question
	err := e.db.Table("question").Where("uuid = ?", uuid).First(&question).Error
	if err != nil {
		return nil, fmt.Errorf(constan.IDNotFound)
	}
	return &question, nil
}

func (e *QuestionRepoImpl) UpdateByUUID(uuid uuid.UUID, question model.Question) (*model.Question, error) {
	var output model.Question
	err := e.db.Table("question").Where("uuid = ?", uuid).First(&output).Update(&question).Error
	if err != nil {
		fmt.Printf("[QuestionRepoImpl.UpdateByUUID] error execute query %v \n", err)
		return nil, fmt.Errorf(constan.FailedUpdateData)
	}
	return &output, nil
}

func (e *QuestionRepoImpl) DeleteByUUID(uuid uuid.UUID) error {
	var question model.Question
	err := e.db.Table("question").Where("uuid = ?", uuid).Delete(question).Error
	if err != nil {
		fmt.Printf("[QuestionRepoImpl.DeleteByUUID] error execute query %v \n", err)
		return fmt.Errorf(constan.FailedInternalServerError)
	}
	return nil
}