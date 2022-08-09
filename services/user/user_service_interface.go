package services

import "github.com/dakasakti/todolist-web/entities"

type UserService interface {
	Register(data entities.UserRequest) error
	Login(data entities.LoginRequest) (string, error)
	GetProfile(user_id uint) (entities.User, error)
}
