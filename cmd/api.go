package main

import (
	"api/cmd/api/configs"
	"api/cmd/api/controllers"
	"api/cmd/api/routers"
	"api/cmd/api/services/db"
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	var cfg configs.Config
	var err error
	r := gin.Default()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	cfg.DB = *mysql.NewConfig()
	cfg.DB.ParseTime = true

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

	db, err := db.OpenDB(&cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	gc := &controllers.GeneralController{
		Database: db,
		Config:   cfg,
		Logger:   logger,
	}

	routers.InitRoutes(r, gc)

	log.Fatal(r.Run())
}
