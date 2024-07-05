package main

import (
	"b3-insert/db"
	"b3-insert/services"
	"fmt"
	"log"
	"sync"
	"time"

	"b3-insert/config"
	"b3-insert/utils"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	dirPath := utils.GetEnv("DIRECTORY_PATH", "")
	if dirPath == "" {
		log.Fatal("DIRECTORY_PATH não definido no arquivo .env")
	}

	databaseURL := utils.GetEnv("DATABASE_URL", "")
	conn, err := db.Connect(databaseURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer conn.Close()

	err = db.Migrate(conn)
	if err != nil {
		log.Fatalf("Erro ao migrar banco de dados: %v", err)
	}

	err = db.ClearTable(conn)
	if err != nil {
		log.Fatalf("Erro ao limpar a tabela: %v", err)
	}

	files, err := utils.GetFilesInDirectory(dirPath)
	if err != nil {
		log.Fatalf("Erro ao obter lista de arquivos: %v", err)
	}

	maxWorkers := utils.Atoi(utils.GetEnv("MAX_WORKERS", "10"))
	semaphore := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	startTime := time.Now()

	for _, file := range files {
		filePath := file
		fmt.Printf("Processando arquivo: %s\n", filePath)

		wg.Add(1)
		semaphore <- struct{}{}

		go func(filePath string) {
			defer wg.Done()
			defer func() { <-semaphore }()
			services.ProcessFile(conn, filePath)
		}(filePath)
	}

	wg.Wait()
	elapsed := time.Since(startTime)
	fmt.Printf("Processamento completo. Tempo total: %s\n", elapsed.Round(time.Second))
}
