package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID          int64          `gorm:"primaryKey" json:"id" deepcopier:"ID"`
	Name        string         `json:"name" deepcopier:"Name"`
	Type        string         `json:"type" deepcopier:"Type"`
	Description string         `json:"description" deepcopier:"Description"`
	UserID      int64          `json:"user_id" deepcopier:"UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Account) TableName() string {
	return "accounts"
}
