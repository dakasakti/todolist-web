package repositories

import "github.com/dakasakti/postingan/entities"

type PostModel interface {
	Insert(data entities.Post) error
	Gets() ([]entities.Post, error)
	Get(id uint) (entities.Post, error)
	Update(id uint, data entities.Post) error
	Delete(id uint) error
}
