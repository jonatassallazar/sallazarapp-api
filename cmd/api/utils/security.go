package utils

import "golang.org/x/crypto/bcrypt"

// SecurePassword faz o hash da string passada para a função
//
// retornando a versão codificada da string ou um erro
func SecurePassword(p string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		return "", err
	}
	return string(hp), err
}

// VerifyPassword verifica se a senha passada bate com o hash atual
//
// retorna erro caso não seja a mesma senha
func VerifyPassword(hp, p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hp), []byte(p))
	if err != nil {
		return err
	}
	return nil
}
