package utils

import (
	"fmt"
	"strings"
)

// isBlank recebe o valor do campo no 1º parâmetro e o nome do campo como 2º parâmetro
//
// Se o campo estiver em branco, será retornado um erro com a mensagem informando o campo em branco
func IsBlank(s string, fn string) error {
	if strings.TrimSpace(s) == "" {
		return fmt.Errorf("o campo %s não pode estar em branco", fn)
	}
	return nil
}
