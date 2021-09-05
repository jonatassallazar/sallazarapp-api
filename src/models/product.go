package models

import (
	"math/big"
	"time"
)

type Product struct {
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Unity     string    `json:"unity"`
	Weight    float64   `json:"weight"`
	PhotoUrl  string    `json:"photo_url"`
	Supplier  string    `json:"supplier"`
	Cost      big.Float `json:"cost"`
	Price     big.Float `json:"price"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
