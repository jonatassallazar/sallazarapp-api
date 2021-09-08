package main

import (
	"api/cmd/api/configs"
	"api/cmd/api/controllers"
	"api/cmd/api/routers"
	"api/cmd/api/services/db"
	"log"
	"os"
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

	configs.LoadEnvironment(&cfg)

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
