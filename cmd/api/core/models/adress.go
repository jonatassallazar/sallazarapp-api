package models

type Address struct {
	CEP           int    `json:"cep,omitempty"`
	Address       string `json:"address,omitempty"`
	Number        string `json:"number,omitempty"`
	Complement    string `json:"complement,omitempty"`
	Neighbourhood string `json:"neighbourhood,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
}
