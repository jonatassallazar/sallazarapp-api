package models

import "time"

type User struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhotoUrl    string    `json:"photo_url"`
	Password    string    `json:"password"`
	AccessLevel string    `json:"access_level"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
