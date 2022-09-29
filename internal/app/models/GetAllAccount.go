package models

// GetAllAccountReq ...
type GetAllAccountReq struct {
	UserID int64 `json:"user_id"`
	Page   int   `json:"page"`
	Limit  int   `json:"limit"`
}

type GetAllAccountRes struct {
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	Limit    int         `json:"limit"`
	LastPage int         `json:"last_page"`
	Data     interface{} `json:"data"`
}
