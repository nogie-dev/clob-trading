package engine

import "github.com/nogie-dev/clob-trading/internal/models"

type Matcher struct {
	books map[string]*models.OrderBook
}

func NewMatcher() *Matcher {
	return &Matcher{
		books: make(map[string]*models.OrderBook),
	}
}
