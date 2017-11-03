package repository

import "github.com/pubgslotsExample/src/model"

type UserRepository interface {
	GetUser(ID int64) (*model.User, error)
	DeleteUser(ID int64) error
	CreateUser(username string) (int64, error)
	Close() error
}
