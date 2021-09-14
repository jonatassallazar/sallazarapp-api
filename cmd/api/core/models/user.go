package models

import (
	"api/cmd/api/utils"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhotoUrl    string    `json:"photo_url,omitempty"`
	Password    string    `json:"password,omitempty"`
	AccessLevel string    `json:"access_level,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)

	if step == "create" {
		hPass, err := utils.SecurePassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hPass)
	}

	return nil
}

func (u *User) validate(step string) error {

	if step == "create" {
		if err := utils.IsBlank(u.Password, "password"); err != nil {
			return err
		}
		if err := utils.IsBlank(u.Name, "name"); err != nil {
			return err
		}
		if err := utils.IsBlank(u.Email, "email"); err != nil {
			return err
		}
	} else {
		if err := utils.IsBlank(u.Email, "email"); err != nil {
			return nil
		} else {
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("o e-mail inserido é inválido")
			}
		}
	}

	return nil
}

// Prepare vai chamar os métodos para validar e formatar o user recebido
func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}

	if err := u.format(step); err != nil {
		return err
	}
	return nil
}
