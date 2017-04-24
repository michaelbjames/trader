package analysis

import (
	"fmt"
	"local/trader/models"
)

type Change int

const (
	Negative = iota - 1
	None
	Positive
)

func ChangeBasedCorrelation(days []models.MarketClose) {
	actionableDays := days[1:]
	changes := make([]Change, len(actionableDays))
	for dayIndex, day := range actionableDays {
		priorDay := days[dayIndex]
		change := day.Price - priorDay.Price
		if change < 0 {
			changes[dayIndex] = Negative
		} else if change == 0 {
			changes[dayIndex] = None
		} else {
			changes[dayIndex] = Positive
		}
	}
	var flips int
	var successes int
	for changeIndex, change := range changes[1:] {
		priorChange := changes[changeIndex]
		if priorChange != change {
			flips++
		}
		if priorChange == change {
			successes++
		}
	}
	fmt.Printf("size: %d; flips: %d; correct implication: %d;",
		len(changes[1:]),
		flips,
		successes)
}
