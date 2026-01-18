package engine

import (
	"container/heap"
	"fmt"
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
	return models.MakerOrder{
		// OrderID: generateID(),
		Ticker:    req.Ticker,
		UserID:    req.UserID,
		OrderType: req.OrderType,
		Position:  req.Position,
		Price:     req.Price,
		Amount:    req.Amount,
		Status:    models.Pending,
		Timestamp: time.Now(),
	}
}

func (ob *OrderBook) AddOrder(order *models.MakerOrder) {
	var levels map[float64]*util.PriceLevel
	var h heap.Interface
	switch order.Position {
	case models.Bid:
		levels, h = ob.Bids, &ob.bidLevels
	case models.Ask:
		levels, h = ob.Asks, &ob.askLevels
	default:
		return
	}

	lvl, ok := levels[order.Price]
	if !ok {
		lvl = &util.PriceLevel{Price: order.Price, Queue: util.NewQueue()}
		levels[order.Price] = lvl
		heap.Push(h, lvl)
	}
	lvl.TotalAmount += order.Amount
	lvl.Queue.Push(order)
}

func (ob *OrderBook) PrintOrderBook() {
	bidHeap := append(util.MaxPriceHeap(nil), ob.bidLevels...)
	heap.Init(&bidHeap)
	for bidHeap.Len() > 0 {
		lvl := heap.Pop(&bidHeap).(*util.PriceLevel)
		fmt.Printf("BID price=%.4f total=%.4f\n", lvl.Price, lvl.TotalAmount)
	}

	askHeap := append(util.MinPriceHeap(nil), ob.askLevels...)
	heap.Init(&askHeap)
	for askHeap.Len() > 0 {
		lvl := heap.Pop(&askHeap).(*util.PriceLevel)
		fmt.Printf("ASK price=%.4f total=%.4f\n", lvl.Price, lvl.TotalAmount)
	}
}
