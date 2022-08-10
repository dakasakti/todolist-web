package repositories

import "github.com/dakasakti/todolist-web/entities"

type UserModel interface {
	Insert(data entities.User) error
	Login(email string) (entities.User, error)
	Get(user_id uint) (entities.User, error)
}
