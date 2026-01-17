package engine

import (
	"fmt"
	"sort"
	"time"

	"github.com/nogie-dev/clob-trading/internal/models"
)

type OrderBook struct {
	Bids   map[float64][]*models.MakerOrder
	Asks   map[float64][]*models.MakerOrder
	Ticker string
}

func NewOrderBook(ticker string) *OrderBook {
	return &OrderBook{
		Ticker: ticker,
		Bids:   make(map[float64][]*models.MakerOrder),
		Asks:   make(map[float64][]*models.MakerOrder),
	}
}

func CreateOrder(req models.RequestOrder) models.MakerOrder {
	order := models.MakerOrder{
		// OrderID:   generateID()
		Ticker:    req.Ticker,
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

func (orderbook *OrderBook) PrintOrderBook() {
	fmt.Printf("=== OrderBook %s ===\n", orderbook.Ticker)

	bidPrices := make([]float64, 0, len(orderbook.Bids))
	for p := range orderbook.Bids {
		bidPrices = append(bidPrices, p)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(bidPrices)))
	for _, p := range bidPrices {
		for _, o := range orderbook.Bids[p] {
			fmt.Printf("[BID] price=%.4f amount=%.4f id=%s\n", o.Price, o.Amount, o.OrderID)
		}
	}

	askPrices := make([]float64, 0, len(orderbook.Asks))
	for p := range orderbook.Asks {
		askPrices = append(askPrices, p)
	}
	sort.Float64s(askPrices)
	for _, p := range askPrices {
		for _, o := range orderbook.Asks[p] {
			fmt.Printf("[ASK] price=%.4f amount=%.4f id=%s\n", o.Price, o.Amount, o.OrderID)
		}
	}
}
