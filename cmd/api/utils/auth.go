package utils

import (
	"api/cmd/api/configs"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CreateToken retorna o token JWT para ser usado nas requisições
//
// ou retorna erro em caso de algum problema na criação do token
func CreateToken(userID uint64, userAccess string) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 48).Unix()
	perms["userID"] = userID
	perms["access"] = userAccess

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)

	return token.SignedString([]byte(configs.SecretJWT))
}

// ValidateToken verifica se o token passado na requisição é válido
func ValidateToken(c *gin.Context) error {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

// extractToken verifica se o token está em formato válido
func extractToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")

	//separa o token do prefixo Bearer
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// returnVerificationKey verifica se o método de assinatura é o esperado
func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return []byte(configs.SecretJWT), nil
}

// ExtractUserIdAndAccessLevel retorna o ID e nível de acesso do usuário que está salvo no token
func ExtractUserIdAndAccessLevel(c *gin.Context) (uint64, string, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, "", err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if err != nil {
			return 0, "", err
		}

		access := fmt.Sprintf("%s", permissions["access"])
		if err != nil {
			return 0, "", err
		}

		return userId, access, nil
	}

	return 0, "", errors.New("token inválido")
}
