package models

// Brokerage is how various trader executions should look.
type Brokerage interface {
	Buy(interface{}) error
	Sell(interface{}) error
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
	Principal float64
}

type AnalysisBuy int

type AnalysisSell int
