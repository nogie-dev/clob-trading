package engine

import (
	"fmt"
	"log"
	"sort"
	"sync"

	"github.com/nogie-dev/clob-trading/internal/models"
)

type Matcher struct {
	mu    sync.RWMutex
	books map[string]*OrderBook
}

func NewMatcher() *Matcher {
	return &Matcher{
		books: make(map[string]*OrderBook),
	}
}

func (m *Matcher) GetOrderBook(ticker string) *OrderBook {
	if _, ok := m.books[ticker]; !ok {
		log.Fatal("Not Exist OrderBook")
		return nil
	}
	return m.books[ticker]
}

func (m *Matcher) ProcessingOrder(orderbook *OrderBook, order models.MakerOrder) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// book := m.GetOrderBook(order.Ticker)
	switch order.Position {
	case models.Bid:
		orderbook.Bids[order.Price] = append(orderbook.Bids[order.Price], &order)
	case models.Ask:
		orderbook.Asks[order.Price] = append(orderbook.Asks[order.Price], &order)
	default:
	}
}

func (m *Matcher) PrintBook(ticker string) {
	m.mu.RLock()
	book, ok := m.books[ticker]
	m.mu.RUnlock()
	if !ok {
		fmt.Printf("no order book for ticker %s\n", ticker)
		return
	}

	fmt.Printf("=== OrderBook %s ===\n", ticker)

	bidPrices := make([]float64, 0, len(book.Bids))
	for p := range book.Bids {
		bidPrices = append(bidPrices, p)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(bidPrices)))
	for _, p := range bidPrices {
		for _, o := range book.Bids[p] {
			fmt.Printf("[BID] price=%.4f amount=%.4f id=%s\n", o.Price, o.Amount, o.OrderID)
		}
	}

	askPrices := make([]float64, 0, len(book.Asks))
	for p := range book.Asks {
		askPrices = append(askPrices, p)
	}
	sort.Float64s(askPrices)
	for _, p := range askPrices {
		for _, o := range book.Asks[p] {
			fmt.Printf("[ASK] price=%.4f amount=%.4f id=%s\n", o.Price, o.Amount, o.OrderID)
		}
	}
}
