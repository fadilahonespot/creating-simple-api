package model

import (
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)

type Question struct {
	gorm.Model
	UUID      uuid.UUID `json:"uuid"`
	Question  string    `json:"question"`
	CreatedBy string    `json:"created_by"`
	UpdateBy  string    `json:"update_by"`
	IsActive  bool      `json:"is_active"`
}

func (e *Question) TableName() string {
	return "question"
}
