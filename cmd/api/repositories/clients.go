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
	stmtClient, err := c.db.Prepare(
		"INSERT INTO clients (name, status, email, phone, gender, birthday, owner_id) values (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmtClient.Close()

	resClient, err := stmtClient.Exec(client.Name, client.Status, client.Email, client.Phone, client.Gender, client.Birthday, userID)
	if err != nil {
		return 0, err
	}

	newClientID, err := resClient.LastInsertId()
	if err != nil {
		return 0, err
	}

	var address models.Address

	stmtAddress, err := c.db.Prepare(
		"INSERT INTO address (cep, number, complement, neighbourhood, city, state, client_id) values (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmtAddress.Close()

	_, err = stmtAddress.Exec(address.CEP, address.Number, address.Complement, address.Neighbourhood, address.City, address.State, newClientID)
	if err != nil {
		return 0, err
	}

	return uint64(newClientID), nil
}

// GetClientsByUserID busca no banco todos os clientes pelo ID do usuário logado
func (c Clients) GetClientsByUserID(userID uint64) ([]models.Client, error) {
	var clients []models.Client

	rows, err := c.db.Query(
		"SELECT clients.id, clients.name, clients.status, clients.email, clients.phone, clients.gender, clients.birthday, clients.owner_id, clients.updated_at, clients.created_at, address.cep, address.number, address.complement, address.neighbourhood, address.city, address.state FROM clients WHERE owner_id = ? LEFT JOIN address ON clients.id = address.client_id", userID)
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
			&client.FullAddress.CEP,
			&client.FullAddress.Number,
			&client.FullAddress.Complement,
			&client.FullAddress.Neighbourhood,
			&client.FullAddress.City,
			&client.FullAddress.State,
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
		"SELECT clients.id, clients.name, clients.status, clients.email, clients.phone, clients.gender, clients.birthday, clients.owner_id, clients.updated_at, clients.created_at, address.cep, address.number, address.complement, address.neighbourhood, address.city, address.state FROM clients WHERE owner_id = ? AND id = ? LEFT JOIN address ON clients.id = address.client_id",
		userID, clientID).Scan(
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
		&client.FullAddress.CEP,
		&client.FullAddress.Number,
		&client.FullAddress.Complement,
		&client.FullAddress.Neighbourhood,
		&client.FullAddress.City,
		&client.FullAddress.State); err != nil {
		return models.Client{}, err
	}

	return client, nil
}

// UpdateClientByID atualiza um cliente no banco pelo ID fornecido
func (c Clients) UpdateClientByID(userID, clientID uint64, client models.Client) error {
	stmtClient, err := c.db.Prepare(
		"UPDATE clients SET name = ?,status = ?, email = ?, phone = ?, gender = ?, birthday = ?, updated_at = ? WHERE id = ? AND owner_id = ?")
	if err != nil {
		return err
	}
	defer stmtClient.Close()

	_, err = stmtClient.Exec(
		client.Name, client.Status, client.Email, client.Phone,
		client.Gender, client.Birthday, time.Now().UTC(), clientID, userID)
	if err != nil {
		return err
	}

	var address models.Address

	stmtAddress, err := c.db.Prepare(
		"UPDATE address SET cep = ?, number = ?, complement = ?, neighbourhood = ?, city = ?, state = ? WHERE client_id = ?")
	if err != nil {
		return err
	}
	defer stmtAddress.Close()

	_, err = stmtAddress.Exec(address.CEP, address.Number, address.Complement, address.Neighbourhood, address.City, address.State, clientID)
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
