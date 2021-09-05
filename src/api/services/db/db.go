package db

import (
	"api/src/configs"
	"context"
	"database/sql"
	"time"
)

// OpenDB abrea uma conexão com o banco de dados
func OpenDB(cfg *configs.Config) (*sql.DB, error) {
	dsn := cfg.DB.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
