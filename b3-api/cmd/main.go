package main

import (
	"b3-api/db"
	"b3-api/handler"
	"b3-api/repository"
	"b3-api/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	db.InitDB()

	repo := repository.NewTradeRepository(db.DB())
	tradeService := service.NewTradeService(repo)
	tradeHandler := handler.NewHandler(tradeService)

	router := mux.NewRouter()

	router.HandleFunc("/api/aggregated-data/{ticker}", tradeHandler.GetAggregatedData).Methods("GET")

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Servidor iniciado na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
