package models

import (
	"time"

	"gorm.io/gorm"
)

type Finance struct {
	ID              int64          `json:"id"  deepcopier:"ID"`
	Title           string         `json:"title" deepcopier:"Title" validate:"required"`
	AccountID       int64          `json:"account_id" deepcopier:"FinanceAccountID" validate:"required,number"`
	Account         Account        `gorm:"foreignKey:AccountID" json:"account"`
	Amount          float64        `json:"amount" deepcopier:"Amount" validate:"required,number"`
	Description     string         `json:"description" deepcopier:"Description" validate:"required"`
	UserID          int64          `json:"user_id" deepcopier:"UserID" validate:"required,number"`
	Type            string         `json:"type" deepcipier:"type"`
	TransactionDate string         `json:"transaction_date" deepcopier:"TransactionDate" validate:"required"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

func (Finance) TableName() string {
	return "finances"
}

type FinanceReq struct {
	ID              int64   `json:"id"  deepcopier:"ID"`
	Title           string  `json:"title" deepcopier:"Title" validate:"required"`
	AccountID       int64   `json:"account_id" deepcopier:"FinanceAccountID" validate:"required,number"`
	Amount          float64 `json:"amount" deepcopier:"Amount" validate:"required,number"`
	Description     string  `json:"description" deepcopier:"Description" validate:"required"`
	UserID          int64   `json:"user_id" deepcopier:"UserID" validate:"required,number"`
	Type            string  `json:"type" deepcipier:"type" validate:"required,oneof=expense income"`
	TransactionDate string  `json:"transaction_date" deepcopier:"TransactionDate" validate:"required"`
}
