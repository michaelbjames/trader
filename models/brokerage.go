package models

import "math/big"

// Brokerage is how various trader executions should look.
type Brokerage interface {
	Buy(interface{}) error
	Sell(interface{}) error
	Trade(*Action)
}

// BrokerageName is the brokerage we'll execute a trade through.
type BrokerageName string

// Some brokers. Analysis does no live stuff.
const (
	BrokerAnalysis  BrokerageName = "analysis"
	BrokerRobinhood BrokerageName = "robinhood"
	BrokerCoindesk  BrokerageName = "coindesk"
)

type AnalysisInit struct {
	Source    []Event
	Principal *big.Float
}

type AnalysisBuy *big.Float

type AnalysisSell *big.Float
