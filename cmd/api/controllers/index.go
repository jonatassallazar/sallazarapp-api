package controllers

import (
	"api/cmd/api/configs"
	"database/sql"
	"log"
)

// GeneralController é um struct que leva database, logger e variáveis de config pelos controllers
type GeneralController struct {
	Database *sql.DB        // Objeto da Database aberta
	Config   configs.Config // Configurações base da api
	Logger   *log.Logger
}
