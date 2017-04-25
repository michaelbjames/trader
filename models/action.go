package models

// Action holds the data required to perform a trade (or not perform one).
type Action struct {
	Type     ActionType
	Symbol   Symbol
	Quantity float64
}

// ActionType are the kinds of an Action.
type ActionType string

// Instances of ActionTypes.
const (
	ActionTypeBuy  = "BUY"
	ActionTypeSell = "SELL"
)

// Symbol is the name of the thing to buy or sell.
type Symbol string
