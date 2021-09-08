package configs

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

// Config são as configurações da base da app
type Config struct {
	Port int
	Env  string
	DB   mysql.Config
}

var SecretJWT string

// LoadEnvironment inicializa todas as variáveis vindas do ambiente
func LoadEnvironment(cfg *Config) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	flag.IntVar(&cfg.Port, "port", port, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.DB.Addr, "db-addr", os.Getenv("MYSQL_ADDR"), "MySql adress")
	flag.StringVar(&cfg.DB.Net, "db-net", os.Getenv("MYSQL_NET"), "MySql net type")
	flag.StringVar(&cfg.DB.DBName, "db-name", os.Getenv("MYSQL_DATABASE"), "MySql database name")
	flag.StringVar(&cfg.DB.User, "db-username", os.Getenv("MYSQL_USER"), "MySql username")
	flag.StringVar(&cfg.DB.Passwd, "db-password", os.Getenv("MYSQL_PASSWORD"), "MySql password secret")
	flag.StringVar(&SecretJWT, "jwt-secret-key", os.Getenv("JWT_SECRET"), "JWT Secret Key")
	flag.Parse()
}
