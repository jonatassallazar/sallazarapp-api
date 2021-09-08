package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// Generate64Key função standalone que cria uma chave 64 bits
func Generate64Key() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	string64Base := base64.StdEncoding.EncodeToString(key)
	fmt.Println(string64Base)
}
