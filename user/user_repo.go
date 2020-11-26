package user

import "creating-simple-api/model"

type UserRepo interface {
	InserUser(user *model.User) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	FindUserById(id int) (*model.User, error)
}