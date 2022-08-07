package repositories

import (
	"github.com/dakasakti/todolist-web/entities"

	"gorm.io/gorm"
)

type postModel struct {
	Db *gorm.DB
}

func NewPostModel(db *gorm.DB) *postModel {
	return &postModel{Db: db}
}

func (pm *postModel) Insert(data entities.Post) error {
	err := pm.Db.Create(&data)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (pm *postModel) Gets() ([]entities.Post, error) {
	var Posts []entities.Post

	err := pm.Db.Preload("User").Order("status, deadline").Find(&Posts)
	if err.Error != nil {
		return nil, err.Error
	}

	return Posts, nil
}

func (pm *postModel) Get(id uint) (entities.Post, error) {
	var Post entities.Post
	err := pm.Db.Preload("User").First(&Post, id)
	if err.Error != nil {
		return Post, err.Error
	}

	return Post, nil
}

func (pm *postModel) Update(id uint, data entities.Post) error {
	err := pm.Db.Where("id = ?", id).Updates(data)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (pm *postModel) Delete(id uint) error {
	err := pm.Db.Where("id = ?", id).Delete(&entities.Post{})
	if err.Error != nil {
		return err.Error
	}

	return nil
}
