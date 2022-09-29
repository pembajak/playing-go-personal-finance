package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `json:"id" deepcopier:"ID"`
	Email     string         `json:"email" deepcopier:"email" validate:"required,email,min=10,max=100"`
	FullName  string         `json:"full_name" deepcopier:"FullName" validate:"required,min=10,max=100"`
	Phone     string         `json:"phone" deepcopier:"Phone" validate:"required,min=8,max=16"`
	Password  string         `json:"password,omitempty" deepcopier:"Password" validate:"required,min=5,max=20"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
