package models

// Executor is how various trader executions should look.
type Executor interface {
	Buy(interface{}) error
	Sell(interface{}) error
}

// Broker is the brokerage we'll execute a trade through.
type Broker string

// Some brokers. Analysis does no live stuff.
const (
	Analysis  Broker = "Analysis"
	Robinhood Broker = "Robinhood"
	Coindesk  Broker = "Coindesk"
)
