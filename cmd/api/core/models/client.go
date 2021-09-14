package models

import (
	"strings"
	"time"
)

type Client struct {
	ID          uint64    `json:"ID"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Status      string    `json:"status,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Birthday    time.Time `json:"birthday,omitempty"`
	OwnerID     uint64    `json:"owner_id"`
	FullAddress Address   `json:"full_address,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (c *Client) format(step string) error {
	c.Name = strings.TrimSpace(c.Name)
	c.Status = strings.TrimSpace(c.Email)
	c.Phone = strings.TrimSpace(c.Phone)
	c.Gender = strings.TrimSpace(c.Gender)

	return nil
}

// Prepare vai chamar os métodos para validar e formatar o user recebido
func (c *Client) Prepare(step string) error {
	if err := c.format(step); err != nil {
		return err
	}
	return nil
}
