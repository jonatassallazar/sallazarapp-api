package repositories

import (
	"api/cmd/api/core/models"
	"database/sql"
	"time"
)

// Users representa um repositório de usuários
type Users struct {
	db *sql.DB
}

// NewUsersRepo cria um repositório de usuários
func NewUsersRepo(db *sql.DB) *Users {
	return &Users{db}
}

// CreateUser cria um usuário no banco
func (u Users) CreateUser(user models.User) (uint64, error) {
	stmt, err := u.db.Prepare(
		"INSERT INTO users (name, email, password, accesslevel) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Email, user.Password, "admin")
	if err != nil {
		return 0, err
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

// GetUserByID busca no banco um usuário pelo seu ID
func (u Users) GetUserByID(ID uint64) (models.User, error) {
	var user models.User

	if err := u.db.QueryRow(
		"SELECT id, name, email, updated_at, created_at FROM users WHERE id = ?", ID).Scan(
		&user.ID, &user.Name, &user.Email, &user.UpdatedAt, &user.CreatedAt,
	); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// GetUserByEmail busca no banco um usuário pelo e-mail
//
// Uso restrito interno da API para autenticar o login de um usuário
func (u Users) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	if err := u.db.QueryRow(
		"SELECT id, email, password, accesslevel FROM users WHERE email = ?", email).Scan(
		&user.ID, &user.Email, &user.Password, &user.AccessLevel,
	); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// UpdateUserByID atualiza um usuário no banco pelo ID fornecido
func (u Users) UpdateUserByID(ID uint64, user models.User) error {
	stmt, err := u.db.Prepare("UPDATE users SET nome = ?, email = ?, updated_at = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, time.Now().UTC(), ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserByID deleta um usuário no banco pelo ID informado
func (u Users) DeleteUserByID(id uint64) error {
	stmt, err := u.db.Prepare("DELETE * FROM users WHERE id = ?")
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
