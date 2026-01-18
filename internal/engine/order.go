package engine

import (
	"container/heap"
	"time"

	"github.com/nogie-dev/clob-trading/internal/models"
	"github.com/nogie-dev/clob-trading/internal/util"
)

type OrderBook struct {
	Bids      map[float64]*util.PriceLevel
	Asks      map[float64]*util.PriceLevel
	bidLevels util.MaxPriceHeap
	askLevels util.MinPriceHeap
	Ticker    string
}

func NewOrderBook(ticker string) *OrderBook {
	ob := &OrderBook{
		Ticker: ticker,
		Bids:   make(map[float64]*util.PriceLevel),
		Asks:   make(map[float64]*util.PriceLevel),
	}
	heap.Init(&ob.bidLevels)
	heap.Init(&ob.askLevels)
	return ob
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

func (ob *OrderBook) ProcessingOrder(order *models.MakerOrder) {
	switch order.Position {
	case models.Bid:
		ob.addBid(order)
	case models.Ask:
		ob.addAsk(order)
	default:
	}
}

func (ob *OrderBook) addBid(order *models.MakerOrder) {
	lvl, ok := ob.Bids[order.Price]
	if !ok {
		lvl = &util.PriceLevel{
			Price: order.Price,
			Queue: util.NewQueue(),
		}
		ob.Bids[order.Price] = lvl
		heap.Push(&ob.bidLevels, lvl)
	}
	lvl.Queue.Push(order)
}

func (ob *OrderBook) addAsk(order *models.MakerOrder) {
	lvl, ok := ob.Asks[order.Price]
	if !ok {
		lvl = &util.PriceLevel{
			Price: order.Price,
			Queue: util.NewQueue(),
		}
		ob.Asks[order.Price] = lvl
		heap.Push(&ob.bidLevels, lvl)
	}
	lvl.Queue.Push(order)
}
