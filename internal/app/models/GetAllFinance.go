package models

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
