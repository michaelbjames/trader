package gateways

import "local/trader/models"

type analysisBroker struct {
	Initial float64
	Current float64
	Actions []models.Action
}

func NewAnalysisBroker() models.Executor {
	return &analysisBroker{}
}

func (b *analysisBroker) Buy(i interface{}) error {
	return nil
}
