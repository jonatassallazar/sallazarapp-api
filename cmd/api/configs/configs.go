package configs

import (
	"github.com/go-sql-driver/mysql"
)

// Config são as configurações da base da app
type Config struct {
	Port int
	Env  string
	DB   mysql.Config
}
