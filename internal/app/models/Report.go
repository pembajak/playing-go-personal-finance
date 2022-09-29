package models

import "time"

// TotalTransactionRes ...
type TotalTransactionRes struct {
	TransactionDate time.Time `json:"transaction_date" deepcopier:"TransactionDate"`
	Total           float64   `json:"total" deepcopier:"Total"`
	Type            string    `json:"type" deepcopier:"Type"`
}

// TotalTransactionMonthlyRes ...
type TotalTransactionMonthlyRes struct {
	Month    string  `json:"month" deepcopier:"Month"`
	MonthNum string  `json:"month_num" deepcopier:"MonthNum"`
	Year     string  `json:"year" deepcopier:"Year"`
	Total    float64 `json:"total" deepcopier:"Total"`
	Type     string  `json:"type" deepcopier:"Type"`
}
