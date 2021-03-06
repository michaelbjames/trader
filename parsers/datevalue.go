package parsers

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"time"

	"os"

	"local/trader/models"

	"github.com/pkg/errors"
)

type options struct {
	DateFmt  string
	Currency models.Currency
}

func datevalueParse(reader *csv.Reader, options options) ([]models.Event, error) {
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "Bad CSV!")
	}
	closes := []models.Event{}
	for lineNumber, line := range lines {
		mc, err := parseLine(line, options.DateFmt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error on line %d: %s\n", lineNumber+1, err.Error())
		} else {
			mc.Currency = options.Currency
			closes = append(closes, mc)
		}
	}
	return closes, nil
}

func parseLine(csvLine []string, dateFmt string) (models.Event, error) {
	datetimeString := csvLine[0]
	datetime, err := time.Parse(dateFmt, datetimeString)
	if err != nil {
		return models.Event{}, errors.Wrap(err, "unable to parse time in line")
	}
	closePrice, _, err := big.ParseFloat(csvLine[1], 10, 64, big.ToNearestEven)
	if err != nil {
		return models.Event{}, errors.Wrap(err, "unable to parse close price")
	}
	return models.Event{
		Datetime: datetime,
		Price:    *closePrice,
	}, nil
}
