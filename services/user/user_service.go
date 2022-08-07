package services

import (
	"errors"

	"github.com/dakasakti/todolist-web/deliveries/middlewares"
	"github.com/dakasakti/todolist-web/entities"
	um "github.com/dakasakti/todolist-web/repositories/user"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	um um.UserModel
}

func NewUserService(um um.UserModel) *userService {
	return &userService{um: um}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (us *userService) Register(data entities.UserRequest) error {
	hash, err := HashPassword(data.Password)
	if err != nil {
		return err
	}

	dataUser := entities.User{
		Fullname: data.Fullname,
		Phone:    data.Phone,
		Email:    data.Email,
		Password: hash,
	}

	err = us.um.Insert(dataUser)
	if err != nil {
		if err.Error() == "Error 1062: Duplicate entry '"+data.Email+"' for key 'email'" {
			return errors.New("email already registered")
		} else if err.Error() == "Error 1062: Duplicate entry '"+data.Phone+"' for key 'phone'" {
			return errors.New("phone number already registered")
		} else {
			return err
		}
	}

	return nil
}

func (us *userService) Login(data entities.LoginRequest) (string, error) {
	dataUser, err := us.um.Login(data.Email)
	if err != nil {
		return "", err
	}

	if !CheckPasswordHash(data.Password, dataUser.Password) {
		return "", errors.New("username or password is wrong")
	}

	result, err := middlewares.CreateToken(dataUser.ID)
	if err != nil {
		return "", err
	}

	return result, nil
}
