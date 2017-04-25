package models

import (
	"time"
)

// Currency symbol.
type Currency string

// Currency instances.
const (
	CurrencyAbstract Currency = "abstract"
	CurrencyBTC      Currency = "BTC"
	CurrencyUSD      Currency = "USD"
	CurrencyLTC      Currency = "LTC"
	CurrencyETH      Currency = "ETH"
)

// Event is the end-of-day value.
type Event struct {
	Datetime time.Time
	Currency Currency
	Price    float64
}
