package controllers

import (
	"api/cmd/api/configs"
	"database/sql"
	"log"
)

type GeneralController struct {
	Database *sql.DB
	Config   configs.Config
	Logger   *log.Logger
}
