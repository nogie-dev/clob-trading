package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/nogie-dev/clob-trading/internal/models"
)

func GenerateOrderID(req models.RequestOrder) string {
	payload := fmt.Sprintf("%s|%s|%s|%s|%.8f|%.8f",
		req.Ticker, req.UserID, req.OrderType, req.Position, req.Price, req.Amount)
	sum := sha256.Sum256([]byte(payload))
	return hex.EncodeToString(sum[:])
}
