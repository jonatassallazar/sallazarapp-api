package repositories

import (
	"api/cmd/api/core/models"
	"database/sql"
)

// Users representa um repositório de usuários
type Users struct {
	db *sql.DB
}

// NewUsersRepo cria um repositório de usuários
func NewUsersRepo(db *sql.DB) *Users {
	return &Users{db}
}

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

func (u Users) GetUserByID(ID uint64) (models.User, error) {
	var user models.User

	if err := u.db.QueryRow(
		"SELECT id, name, email, password, updated_at, created_at FROM users WHERE id = ?", ID).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.UpdatedAt, &user.CreatedAt,
	); err != nil {
		return models.User{}, err
	}

	return user, nil
}

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
