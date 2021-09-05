package main

import (
	"api/src/configs"
	"api/src/db"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

const version = "1.0.0"

func main() {
	var cfg configs.Config
	var err error

	cfg.DB = *mysql.NewConfig()

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
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	/* app := &configs.Application{
		Config: cfg,
		Logger: logger,
	 }*/

	r := gin.Default()

	db, err := db.OpenDB(&cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(r.Run())
}
