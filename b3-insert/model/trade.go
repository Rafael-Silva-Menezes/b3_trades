package model

import (
	"time"
)

type Trade struct {
	ID             string
	Ticker         string
	TradePrice     float64
	TradedQuantity int
	ClosingTime    string
	TradeDate      time.Time
}
