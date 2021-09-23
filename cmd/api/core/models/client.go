package models

import (
	"strings"
	"time"
)

type Client struct {
	ID          uint64
	Name        string
	Email       NullString
	Status      string
	Phone       NullString
	Gender      NullString
	Birthday    NullTime
	OwnerID     uint64
	FullAddress Address
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type ClientJSON struct {
	ID          uint64      `json:"ID"`
	Name        string      `json:"name,omitempty"`
	Email       string      `json:"email,omitempty"`
	Status      string      `json:"status,omitempty"`
	Phone       string      `json:"phone,omitempty"`
	Gender      string      `json:"gender,omitempty"`
	Birthday    time.Time   `json:"birthday,omitempty"`
	OwnerID     uint64      `json:"owner_id"`
	FullAddress AddressJSON `json:"full_address,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
}

func (c *ClientJSON) format(step string) error {
	c.Name = strings.TrimSpace(c.Name)
	c.Email = strings.TrimSpace(c.Email)
	c.Phone = strings.TrimSpace(c.Phone)
	c.Gender = strings.TrimSpace(c.Gender)

	return nil
}

// Prepare vai chamar os métodos para validar e formatar o user recebido
func (c *Client) Prepare(step string) (ClientJSON, error) {
	clientJson, err := c.ClientToJSON()
	if err != nil {
		return ClientJSON{}, err
	}

	if err := clientJson.format(step); err != nil {
		return ClientJSON{}, err
	}

	return clientJson, nil
}

func (c *Client) ClientToJSON() (ClientJSON, error) {
	email := c.Email.CheckNullValue()
	phone := c.Phone.CheckNullValue()
	gender := c.Gender.CheckNullValue()
	birthday := c.Birthday.CheckNullValue()

	cep := c.FullAddress.CEP.CheckNullValue()
	address := c.FullAddress.Address.CheckNullValue()
	number := c.FullAddress.Number.CheckNullValue()
	complement := c.FullAddress.Complement.CheckNullValue()
	neighbourhood := c.FullAddress.Neighbourhood.CheckNullValue()
	city := c.FullAddress.City.CheckNullValue()
	state := c.FullAddress.State.CheckNullValue()

	return ClientJSON{
		ID:       c.ID,
		Name:     c.Name,
		Email:    email,
		Status:   c.Status,
		Phone:    phone,
		Gender:   gender,
		Birthday: birthday,
		OwnerID:  c.OwnerID,
		FullAddress: AddressJSON{
			ID:            c.FullAddress.ID,
			CEP:           cep,
			Address:       address,
			Number:        number,
			Complement:    complement,
			Neighbourhood: neighbourhood,
			City:          city,
			State:         state,
		},
		UpdatedAt: c.UpdatedAt,
		CreatedAt: c.CreatedAt,
	}, nil
}
