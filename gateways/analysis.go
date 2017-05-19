package gateways

import (
	"errors"
	"local/trader/config"
	"local/trader/models"
	"math/big"
)

type analysisBroker struct {
	Initial      *big.Float
	Current      *big.Float
	Holdings     *big.Float
	Actions      []models.Action
	Source       []models.Event
	SourceOffset int
}

func NewAnalysisBroker(conf models.AnalysisInit) models.Brokerage {
	return &analysisBroker{
		Initial:      big.NewFloat(float64(config.Get().Brokerages.Analysis.Initial)),
		Current:      big.NewFloat(float64(config.Get().Brokerages.Analysis.Initial)),
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
	purchaseValue := big.NewFloat(0).Mul(
		&marketValue.Price,
		quantity)
	if purchaseValue.Cmp(b.Current) > 0 {
		return errors.New("Analysis Buy: Insufficient funds.")
	}
	b.Current.Sub(b.Current, purchaseValue)
	b.Holdings.Add(b.Holdings, quantity)
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
	purchaseValue := big.NewFloat(0).Mul(
		&marketValue.Price,
		quantity)
	b.Current.Add(b.Current, purchaseValue)
	b.Holdings.Sub(b.Holdings, quantity)
	b.SourceOffset++
	return nil
}

func (b *analysisBroker) Trade(action *models.Action) {
	if action == nil {
		return
	}
	switch action.Type {
	case models.ActionTypeBuy:
		quantity := models.AnalysisBuy(&action.Quantity)
		b.Buy(quantity)
	case models.ActionTypeSell:
		quantity := models.AnalysisSell(&action.Quantity)
		b.Sell(quantity)
	}
}
