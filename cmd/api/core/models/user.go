package models

import "time"

type User struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhotoUrl    string    `json:"photo_url,omitempty"`
	Password    string    `json:"password,omitempty"`
	AccessLevel string    `json:"access_level,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
