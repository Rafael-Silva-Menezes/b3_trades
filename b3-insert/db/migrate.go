package db

import (
	"database/sql"
	"fmt"
)

func Migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS trades (
		id VARCHAR(36) PRIMARY KEY,
		ticker VARCHAR(50),
		trade_price FLOAT8,
		traded_quantity INT,
		closing_time VARCHAR(50),
		trade_date TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela: %w", err)
	}
	return nil
}

func ClearTable(db *sql.DB) error {
	query := `DELETE FROM trades;`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao limpar a tabela: %w", err)
	}
	return nil
}
