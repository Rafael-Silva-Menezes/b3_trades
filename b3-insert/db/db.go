package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	maxOpenConns := 50
	maxIdleConns := 25

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
