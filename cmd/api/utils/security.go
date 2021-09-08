package utils

import "golang.org/x/crypto/bcrypt"

func SecurePassword(p string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return "", err
	}
	return string(hp), err
}

func VerifyPassword(hp, p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hp), []byte(p))
	if err != nil {
		return err
	}
	return nil
}
