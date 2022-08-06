package entities

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Description string         `json:"description"`
	Name        string         `json:"name"`
	Deadline    string         `json:"deadline" gorm:"date"`
	Status      string         `json:"status"`
	UserID      uint           `json:"-"`
	User        User           `json:"author" gorm:"foreignkey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type PostRequest struct {
	Description string `json:"description" form:"description" validate:"required,min=5,max=255"`
	Name        string `json:"name" form:"name" validate:"required,min=5,max=255"`
	Deadline    string `json:"deadline" form:"deadline" validate:"required"`
}

type PostUpdateRequest struct {
	Description string `json:"description" form:"description" validate:"omitempty,min=5,max=255"`
	Name        string `json:"name" form:"name" validate:"omitempty,min=5,max=255"`
	Deadline    string `json:"deadline" form:"deadline" validate:"omitempty"`
}
