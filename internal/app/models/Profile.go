package models

import "time"

type Profile struct {
	ID        int64     `json:"id" deepcopier:"ID"`
	Email     string    `json:"email" deepcopier:"email"`
	FullName  string    `json:"full_name" deepcopier:"FullName"`
	Phone     string    `json:"phone" deepcopier:"Phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
