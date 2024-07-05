package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conectado ao banco de dados PostgreSQL...")

	// Criar índices no banco de dados, se necessário
	_, err = db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_ticker ON trades(ticker);
		-- Adicione outros índices conforme necessário para trade_price, traded_quantity, etc.
	`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Índices criados no banco de dados...")
}

func DB() *sql.DB {
	return db
}
