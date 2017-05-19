package updaters

import (
	"local/trader/models"
	"math/big"
)

const daysToHold int = 2

func NewManicMomentum(initialAmount *big.Float) *manicMomentum {
	return &manicMomentum{
		PastCloses:     []*big.Float{},
		AvailableFunds: initialAmount,
		HeldQuantity:   big.NewFloat(0),
	}
}

type manicMomentum struct {
	PastCloses     []*big.Float
	AvailableFunds *big.Float
	HeldQuantity   *big.Float
}

func (mm *manicMomentum) SaveEvent(e models.Event) (models.History, error) {
	if len(mm.PastCloses) < daysToHold {
		mm.PastCloses = append(mm.PastCloses, &e.Price)
		return mm, nil
	}
	mm.PastCloses = append(mm.PastCloses[1:], &e.Price)
	return mm, nil
}

// What if the action fails? Should there be a revert-state function?
func (mm *manicMomentum) TakeAction() (*models.Action, error) {
	if len(mm.PastCloses) < daysToHold {
		return nil, nil
	}
	delta := mm.PastCloses[0].Cmp(mm.PastCloses[1])
	if delta > 0 {
		buyAmt := new(big.Float).Quo(mm.PastCloses[1], mm.AvailableFunds)
		mm.HeldQuantity.Add(mm.HeldQuantity, buyAmt)
		mm.AvailableFunds = new(big.Float)
		return &models.Action{
			Type:     models.ActionTypeBuy,
			Quantity: *buyAmt,
		}, nil
	}
	if delta < 0 {
		amtWorth := new(big.Float).Mul(mm.HeldQuantity, mm.PastCloses[1])
		mm.AvailableFunds.Add(mm.AvailableFunds, amtWorth)
		return &models.Action{
			Type:     models.ActionTypeSell,
			Quantity: *mm.HeldQuantity,
		}, nil
	}
	return nil, nil
}
