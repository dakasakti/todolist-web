package repositories

import "github.com/dakasakti/postingan/entities"

type UserModel interface {
	Insert(data entities.User) error
	Login(email string) (entities.User, error)
}
