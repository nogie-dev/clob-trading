package main

import (
	"github.com/nogie-dev/clob-trading/internal/engine"
	"github.com/nogie-dev/clob-trading/internal/testdata"
)

func main() {
	// matcher := engine.NewMatcher()
	orderbook := engine.NewOrderBook("BTC-USD")

	for _, req := range testdata.SampleOrders {
		order := engine.CreateOrder(req)
		orderbook.ProcessingOrder(&order)
	}
}
