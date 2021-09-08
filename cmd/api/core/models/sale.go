package models

import (
	"math/big"
	"time"
)

type Sale struct {
	Status               string         `json:"status"`
	Number               string         `json:"number"`
	Date                 string         `json:"date"`
	Shipping             big.Float      `json:"shipping"`
	Discount             big.Float      `json:"discount"`
	Tax                  big.Float      `json:"tax"`
	Subtotal             big.Float      `json:"subtotal"`
	Total                big.Float      `json:"total"`
	Products             []SoldItens    `json:"sold_itens"`
	PaymentMethod        string         `json:"payment_method"`
	InstallmentsQuantity int            `json:"installments_quantity"`
	Installments         []Installments `json:"installments"`
	Client               Client         `json:"client"`
	Observation          string         `json:"observation"`
	UpdatedAt            time.Time      `json:"updated_at"`
	CreatedAt            time.Time      `json:"created_at"`
}

type SoldItens struct {
	Product    Product   `json:"product"`
	Quantity   float64   `json:"quantity"`
	TotalValue big.Float `json:"total_value"`
	UpdatedAt            time.Time      `json:"updated_at"`
	CreatedAt            time.Time      `json:"created_at"`
}

type Installments struct {
	Number int       `json:"number"`
	Value  big.Float `json:"value"`
	Date   time.Time `json:"date"`
	UpdatedAt            time.Time      `json:"updated_at"`
	CreatedAt            time.Time      `json:"created_at"`
}
