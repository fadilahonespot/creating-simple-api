package repo

import (
	"creating-simple-api/constan"
	"creating-simple-api/model"
	"creating-simple-api/user"
	"fmt"

	"github.com/jinzhu/gorm"
)


type UserRepoImpl struct {
	db *gorm.DB
}

func CreateUserRepo(db *gorm.DB) user.UserRepo {
	return &UserRepoImpl{db}
}

func (e *UserRepoImpl) InserUser(user *model.User) (*model.User, error) {
	err := e.db.Save(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.InserUser] error execute query %v \n", err)
		return nil, fmt.Errorf(constan.FailedInsertData)
	}
	return user, nil
}

func (e *UserRepoImpl) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := e.db.Table("user").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf(constan.EmailNotFound)
	}
	return &user, nil
}

func (e *UserRepoImpl) FindUserById(id int) (*model.User, error) {
	var user model.User
	err := e.db.Table("user").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf(constan.IDNotFound)
	}
	return &user, nil
}