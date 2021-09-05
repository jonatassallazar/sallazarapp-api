package configs

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

// configurações da base da app
type Config struct {
	Port int
	Env  string
	DB   mysql.Config
}

// base da app
type Application struct {
	Config Config
	Logger *log.Logger
}
