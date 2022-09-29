package models

import (
	"time"

	"gorm.io/gorm"
)

type Finance struct {
	ID              int64          `json:"id"  deepcopier:"ID"`
	Title           string         `json:"title" deepcopier:"Title"`
	AccountID       int64          `json:"account_id" deepcopier:"FinanceAccountID"`
	Account         Account        `gorm:"foreignKey:AccountID" json:"account"`
	Amount          float64        `json:"amount" deepcopier:"Amount"`
	Description     string         `json:"description" deepcopier:"Description"`
	UserID          int64          `json:"user_id" deepcopier:"UserID"`
	Type            string         `json:"type" deepcipier:"type"`
	TransactionDate string         `json:"transaction_date" deepcopier:"TransactionDate"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

func (Finance) TableName() string {
	return "finances"
}

type GetAllFinanceReq struct {
	UserID    int64  `json:"user_id" deepcopier:"UserID"`
	Page      int    `json:"page" deepcopier:"Page"`
	Limit     int    `json:"limit" deepcopier:"Limit"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Type      string `json:"type"`
}

// GetAllFinanceRes ...
type GetAllFinanceRes struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	Limit    int         `json:"limit"`
	LastPage int         `json:"last_page"`
	Data     interface{} `json:"data"`
}

// GetTotalTransaction ...
type GetTotalTransaction struct {
	UserID    int64  `json:"user_id" deepcopier:"UserID"`
	StartDate string `json:"start_date" deepcopier:"StartDate"`
	EndDate   string `json:"end_date" deepcopier:"EndDate"`
	Type      string `json:"type" deepcopier:"Type"`
}
