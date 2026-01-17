package engine

import (
	"time"

	"github.com/nogie-dev/clob-trading/internal/models"
)

func NewOrderBook(ticker string) *models.OrderBook {
	return &models.OrderBook{
		Ticker: ticker,
		Bids:   make(map[float64][]*models.MakerOrder),
		Asks:   make(map[float64][]*models.MakerOrder),
	}
}

func CreateOrder(req models.RequestOrder) models.MakerOrder {
	order := models.MakerOrder{
		// OrderID:   generateID()
		UserID:    req.UserID,
		OrderType: req.OrderType,
		Position:  req.Position,
		Price:     req.Price,
		Amount:    req.Amount,
		Status:    models.Pending,
		Timestamp: time.Now(),
	}
	return order
}
