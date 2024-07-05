package handler

import (
	"b3-api/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *service.TradeService
}

func NewHandler(service *service.TradeService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetAggregatedData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticker := vars["ticker"]
	date := r.URL.Query().Get("date")

	aggregatedData, err := h.Service.GetAggregatedData(ticker, date)
	if err != nil {
		log.Println("Erro ao obter dados agregados:", err)
		http.Error(w, "Erro ao obter dados agregados", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aggregatedData)
}
