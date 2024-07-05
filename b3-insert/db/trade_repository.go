package db

import (
	"b3-insert/model"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

func InsertBatch(db *sql.DB, batch []model.Trade) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %v", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(pq.CopyIn("trades", "id", "ticker", "trade_price", "traded_quantity", "closing_time", "trade_date"))
	if err != nil {
		return fmt.Errorf("erro ao preparar declaração: %v", err)
	}
	defer stmt.Close()

	for _, trade := range batch {
		_, err = stmt.Exec(trade.ID, trade.Ticker, trade.TradePrice, trade.TradedQuantity, trade.ClosingTime, trade.TradeDate)
		if err != nil {
			pgErr, ok := err.(*pq.Error)
			if !ok {
				return fmt.Errorf("erro ao executar declaração: %v", err)
			}

			if pgErr.Code.Name() == "unique_violation" {
				continue
			}

			return fmt.Errorf("erro ao executar declaração: %v", err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("erro ao finalizar declaração: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("erro ao fazer commit da transação: %v", err)
	}

	return nil
}
