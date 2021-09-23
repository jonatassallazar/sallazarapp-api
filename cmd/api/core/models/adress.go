package models

type Address struct {
	ID            uint64
	CEP           NullInt64
	Address       NullString
	Number        NullString
	Complement    NullString
	Neighbourhood NullString
	City          NullString
	State         NullString
}

type AddressJSON struct {
	ID            uint64
	CEP           int64  `json:"cep,omitempty"`
	Address       string `json:"address,omitempty"`
	Number        string `json:"number,omitempty"`
	Complement    string `json:"complement,omitempty"`
	Neighbourhood string `json:"neighbourhood,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
}
