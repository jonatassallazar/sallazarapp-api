package models

import "time"

type Client struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Status      string    `json:"status"`
	Phone       string    `json:"phone"`
	Gender      string    `json:"gender"`
	Birthday    string    `json:"birthday"`
	FullAddress Address   `json:"full_address"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
