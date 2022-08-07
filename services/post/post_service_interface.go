package services

import "github.com/dakasakti/todolist-web/entities"

type PostService interface {
	Register(user_id uint, data entities.PostRequest) error
	GetAll() ([]entities.Post, error)
	CheckParamId(id string) (uint, error)
	GetById(id uint) (entities.Post, error)
	CheckUser(id uint, user_id uint) (entities.Post, error)
	UpdateById(id uint, data entities.PostUpdateRequest) error
	UpdateMarkById(id uint) error
	DeleteById(id uint) error
}
