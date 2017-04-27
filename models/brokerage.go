package models

// Brokerage is how various trader executions should look.
type Brokerage interface {
	Buy(interface{}) error
	Sell(interface{}) error
}

// Broker is the brokerage we'll execute a trade through.
type Broker string

// Some brokers. Analysis does no live stuff.
const (
	BrokerAnalysis  Broker = "Analysis"
	BrokerRobinhood Broker = "Robinhood"
	BrokerCoindesk  Broker = "Coindesk"
)
