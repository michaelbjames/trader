package models

import (
	"time"
)

type Currency string

const (
	CurrencyBTC Currency = "BTC"
	CurrencyUSD Currency = "USD"
	CurrencyLTC Currency = "LTC"
	CurrencyETH Currency = "ETH"
)

type MarketClose struct {
	Datetime time.Time
	Currency Currency
	Price    float64
}
