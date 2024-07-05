package service

import (
	"b3-api/model"
	"b3-api/repository"
)

type TradeService struct {
	Repo *repository.TradeRepository
}

func NewTradeService(repo *repository.TradeRepository) *TradeService {
	return &TradeService{
		Repo: repo,
	}
}

func (s *TradeService) GetAggregatedData(ticker, date string) (model.AggregatedData, error) {
	var aggregatedData model.AggregatedData

	maxRangeValue, err := s.Repo.GetMaxTradePrice(ticker, date)
	if err != nil {
		return model.AggregatedData{}, err
	}

	maxDailyVolume, err := s.Repo.GetMaxDailyVolume(ticker, date)
	if err != nil {
		return model.AggregatedData{}, err
	}

	aggregatedData.Ticker = ticker
	aggregatedData.MaxRangeValue = maxRangeValue
	aggregatedData.MaxDailyVolume = maxDailyVolume

	return aggregatedData, nil
}
