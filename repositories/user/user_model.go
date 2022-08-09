package repositories

import (
	"errors"

	"github.com/dakasakti/todolist-web/entities"

	"gorm.io/gorm"
)

type userModel struct {
	Db *gorm.DB
}

func NewUserModel(db *gorm.DB) *userModel {
	return &userModel{Db: db}
}

func (um *userModel) Insert(data entities.User) error {
	err := um.Db.Create(&data)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (um *userModel) Login(email string) (entities.User, error) {
	var result entities.User
	err := um.Db.Where("email = ?", email).First(&result)
	if err.Error != nil {
		return result, errors.New("username or password is wrong")
	}

	return result, nil
}

func (um *userModel) Get(user_id uint) (entities.User, error) {
	var result entities.User
	err := um.Db.Where("id = ?", user_id).First(&result)
	if err.Error != nil {
		return result, err.Error
	}

	return result, nil
}
