package gateways

import (
	"errors"
	"local/trader/config"
	"local/trader/models"
)

type analysisBroker struct {
	Initial      float64
	Current      float64
	Holdings     int
	Actions      []models.Action
	Source       []models.Event
	SourceOffset int
}

func NewAnalysisBroker(conf models.AnalysisInit) models.Brokerage {
	return &analysisBroker{
		Initial:      float64(config.Get().Brokerages.Analysis.Initial),
		Current:      float64(config.Get().Brokerages.Analysis.Initial),
		Source:       conf.Source,
		SourceOffset: 0,
	}
}

// Buy executes a buy of a quantity of our fictional asset.
// TODO:2017-04-26:This should have a lock so that I can simulate network delays.
func (b *analysisBroker) Buy(i interface{}) error {
	if b.SourceOffset >= len(b.Source) {
		return errors.New("Analysis Buy: End of source material.")
	}
	quantity, ok := i.(models.AnalysisBuy)
	if !ok {
		return errors.New("Analysis Buy: Invalid purchase action.")
	}
	marketValue := b.Source[b.SourceOffset]
	purchaseValue := marketValue.Price * float64(quantity)
	if purchaseValue > b.Current {
		return errors.New("Analysis Buy: Insufficient funds.")
	}
	b.Current -= purchaseValue
	b.Holdings += int(quantity)
	b.SourceOffset++
	return nil
}

func (b *analysisBroker) Sell(i interface{}) error {
	if b.SourceOffset >= len(b.Source) {
		return errors.New("Analysis Sell: End of source material.")
	}
	quantity, ok := i.(models.AnalysisSell)
	if !ok {
		return errors.New("Analysis Sell: Invalid purchase action.")
	}
	marketValue := b.Source[b.SourceOffset]
	purchaseValue := marketValue.Price * float64(quantity)
	b.Current += purchaseValue
	b.Holdings -= int(quantity)
	b.SourceOffset++
	return nil
}
