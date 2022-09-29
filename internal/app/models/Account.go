package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID          int64          `gorm:"primaryKey" json:"id" deepcopier:"ID"`
	Name        string         `json:"name" deepcopier:"Name" validate:"required"`
	Type        string         `json:"type" deepcopier:"Type" validate:"required"`
	Description string         `json:"description" deepcopier:"Description" validate:"required"`
	UserID      int64          `json:"user_id" deepcopier:"UserID" validate:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
}

type AccountReq struct {
	ID          int64  `gorm:"primaryKey" json:"id" deepcopier:"ID"`
	Name        string `json:"name" deepcopier:"Name" validate:"required"`
	Type        string `json:"type" deepcopier:"Type" validate:"required"`
	Description string `json:"description" deepcopier:"Description" validate:"required"`
}

func (Account) TableName() string {
	return "accounts"
}
