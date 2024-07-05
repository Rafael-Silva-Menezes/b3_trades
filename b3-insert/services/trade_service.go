package services

import (
	"b3-insert/db"
	"b3-insert/model"
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"
	"time"

	"b3-insert/utils"
	"github.com/google/uuid"
)

func ProcessFile(conn *sql.DB, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Erro ao abrir o arquivo %s: %v", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	scanner.Split(bufio.ScanLines)

	if scanner.Scan() {
		_ = scanner.Text()
	}

	batchSize := 10000
	batch := make([]model.Trade, 0, batchSize)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")

		if len(fields) < 9 {
			continue
		}

		trade := model.Trade{
			ID:             uuid.New().String(),
			Ticker:         fields[1],
			TradePrice:     utils.Atof(fields[3]),
			TradedQuantity: utils.Atoi(fields[4]),
			ClosingTime:    fields[5],
		}

		tradeDate, err := time.Parse("2006-01-02", fields[8])
		if err != nil {
			log.Printf("Erro ao converter data: %v", err)
			continue
		}
		trade.TradeDate = tradeDate

		batch = append(batch, trade)

		if len(batch) >= batchSize {
			err := db.InsertBatch(conn, batch)
			if err != nil {
				log.Printf("Erro ao inserir batch no banco de dados: %v", err)
			}
			batch = make([]model.Trade, 0, batchSize)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Erro durante a leitura do arquivo %s: %v", filePath, err)
		return
	}

	if len(batch) > 0 {
		err := db.InsertBatch(conn, batch)
		if err != nil {
			log.Printf("Erro ao inserir batch no banco de dados: %v", err)
		}
	}
}
