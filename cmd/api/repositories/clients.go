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
func (c Clients) CreateClient(client models.Client, userID uint64) (uint64, error) {
	stmt, err := c.db.Prepare(
		"INSERT INTO clients (name, status, email, phone, gender, birthday, owner_id) values (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(client.Name, client.Status, client.Email, client.Phone, client.Gender, client.Birthday, userID)
	if err != nil {
		return 0, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

// GetClientsByUserID busca no banco todos os clientes pelo ID do usuário logado
func (c Clients) GetClientsByUserID(userID uint64) ([]models.Client, error) {
	var clients []models.Client

	rows, err := c.db.Query(
		"SELECT id, name, status, email, phone, gender, birthday, owner_id, updated_at, created_at FROM clients WHERE owner_id = ?", userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var client models.Client

		if err = rows.Scan(
			&client.ID,
			&client.Name,
			&client.Status,
			&client.Email,
			&client.Phone,
			&client.Gender,
			&client.Birthday,
			&client.OwnerID,
			&client.UpdatedAt,
			&client.CreatedAt,
		); err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// GetClientByID busca no banco um cliente pelo seu ID
func (c Clients) GetClientByID(userID, clientID uint64) (models.Client, error) {
	var client models.Client

	if err := c.db.QueryRow(
		"SELECT id, name, status, email, phone, gender, birthday, owner_id, updated_at, created_at FROM clients WHERE owner_id = ? AND id = ?",
		userID, clientID).Scan(&client.ID, &client.Name, &client.Status, &client.Email, &client.Phone, &client.Gender, &client.Birthday,
		&client.OwnerID, &client.UpdatedAt, &client.CreatedAt); err != nil {
		return models.Client{}, err
	}

	return client, nil
}

// UpdateClientByID atualiza um cliente no banco pelo ID fornecido
func (c Clients) UpdateClientByID(userID, clientID uint64, client models.Client) error {
	stmt, err := c.db.Prepare(
		"UPDATE clients SET name = ?,status = ?, email = ?, phone = ?, gender = ?, birthday = ?, updated_at = ? WHERE id = ? AND owner_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		client.Name, client.Status, client.Email, client.Phone,
		client.Gender, client.Birthday, time.Now().UTC(), clientID, userID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteClientByID deleta um cliente no banco pelo ID informado
func (c Clients) DeleteClientByID(clientID, userID uint64) error {
	stmt, err := c.db.Prepare("DELETE * FROM clients WHERE id = ? AND owner_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(clientID, userID)
	if err != nil {
		return err
	}

	return nil
}
