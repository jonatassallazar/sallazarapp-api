package models

import "time"

type Client struct {
	ID          uint64    `json:"ID"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Status      string    `json:"status,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Birthday    string    `json:"birthday,omitempty"`
	FullAddress Address   `json:"full_address,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
