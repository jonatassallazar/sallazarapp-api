package repositories

import (
	"api/cmd/api/core/models"
	"database/sql"
	"time"
)

// Clients representa um repositório de clientes
type Clients struct {
	db *sql.DB
}

// NewClientsRepo cria um repositório de clientes
func NewClientsRepo(db *sql.DB) *Clients {
	return &Clients{db}
}

// CreateClient cria um cliente no banco
func (c Clients) CreateClient(client models.Client) (uint64, error) {
	stmt, err := c.db.Prepare(
		"INSERT INTO clients (name, status, email, phone, gender, birthday) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(client.Name, client.Status, client.Email, client.Phone, client.Gender, client.Birthday)
	if err != nil {
		return 0, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

// GetClientByID busca no banco um cliente pelo ID
func (c Clients) GetClientByID(ID uint64) (models.Client, error) {
	var client models.Client

	if err := c.db.QueryRow(
		"SELECT id, name, status, email, phone, gender, birthday, updated_at, created_at FROM clients WHERE id = ?",
		ID).Scan(&client.ID, &client.Name, &client.Email, &client.UpdatedAt, &client.CreatedAt); err != nil {
		return models.Client{}, err
	}

	return client, nil
}

// UpdateClientByID atualiza um cliente no banco pelo ID fornecido
func (c Clients) UpdateClientByID(ID uint64, client models.Client) error {
	stmt, err := c.db.Prepare(
		"UPDATE clients SET nome = ?,status = ?, email = ?, phone = ?, gender = ?, birthday = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		client.Name, client.Status, client.Email, client.Phone,
		client.Gender, client.Birthday, time.Now().UTC(), ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteClientByID deleta um cliente no banco pelo ID informado
func (c Clients) DeleteClientByID(id uint64) error {
	stmt, err := c.db.Prepare("DELETE * FROM clients WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
