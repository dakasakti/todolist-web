package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"-" gorm:"primarykey"`
	Fullname  string         `json:"fullname" gorm:"not null"`
	Phone     string         `json:"-" gorm:"unique"`
	Email     string         `json:"-" gorm:"unique"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required,min=5,max=25"`
	Phone    string `json:"phone" form:"phone" validate:"required,e164,min=10,max=20"`
	Email    string `json:"email" form:"email" validate:"required,lowercase,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=50"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,lowercase,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=50"`
}
