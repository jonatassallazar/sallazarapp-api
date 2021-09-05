package models

type Address struct {
	CEP           int    `json:"cep"`
	Address       string `json:"address"`
	Number        string `json:"number"`
	Complement    string `json:"complement"`
	Neighbourhood string `json:"neighbourhood"`
	City          string `json:"city"`
	State         string `json:"state"`
}
