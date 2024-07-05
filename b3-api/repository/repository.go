package repository

import (
	"database/sql"
	"log"
	"sync"
)

type TradeRepository struct {
	DB *sql.DB
}

func NewTradeRepository(db *sql.DB) *TradeRepository {
	return &TradeRepository{
		DB: db,
	}
}

func (repo *TradeRepository) GetMaxTradePrice(ticker, date string) (float64, error) {
	var maxTradePrice float64
	query := `
		SELECT MAX(trade_price) 
		FROM trades 
		WHERE ticker = $1
	`
	params := []interface{}{ticker}

	if date != "" {
		query += ` AND trade_date >= $2`
		params = append(params, date)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := repo.DB.QueryRow(query, params...).Scan(&maxTradePrice)
		if err != nil && err != sql.ErrNoRows {
			log.Println("Erro ao consultar max_trade_price:", err)
			maxTradePrice = 0
		}
	}()

	wg.Wait()
	return maxTradePrice, nil
}

func (repo *TradeRepository) GetMaxDailyVolume(ticker, date string) (int, error) {
	var maxDailyVolume int
	query := `
		SELECT SUM(traded_quantity) 
		FROM trades 
		WHERE ticker = $1
	`
	params := []interface{}{ticker}

	if date != "" {
		query += ` AND trade_date >= $2`
		params = append(params, date)
	}
	query += `
		GROUP BY trade_date 
		ORDER BY SUM(traded_quantity) DESC 
		LIMIT 1
	`

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := repo.DB.QueryRow(query, params...).Scan(&maxDailyVolume)
		if err != nil && err != sql.ErrNoRows {
			log.Println("Erro ao consultar max_daily_volume:", err)
			maxDailyVolume = 0
		}
	}()

	wg.Wait()
	return maxDailyVolume, nil
}
