package parsers

import (
	"encoding/csv"
	"fmt"
	"local/trader/models"
	"strconv"
	"time"
)
import "github.com/pkg/errors"

const coindeskDateFmt string = "2006-01-02 15:04:05"

func CoindeskMarketClose(reader *csv.Reader) ([]models.MarketClose, error) {
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "Bad CSV!")
	}
	closes := []models.MarketClose{}
	for lineNumber, line := range lines {
		mc, err := parseLine(line)
		if err != nil {
			fmt.Printf("error on line %d: %s\n", lineNumber+1, err.Error())
		} else {
			closes = append(closes, mc)
		}
	}
	return closes, nil
}

func parseLine(csvLine []string) (models.MarketClose, error) {
	datetimeString := csvLine[0]
	datetime, err := time.Parse(coindeskDateFmt, datetimeString)
	if err != nil {
		return models.MarketClose{}, errors.Wrap(err, "unable to parse time in line")
	}
	closePrice, err := strconv.ParseFloat(csvLine[1], 64)
	if err != nil {
		return models.MarketClose{}, errors.Wrap(err, "unable to parse close price")
	}
	return models.MarketClose{
		Datetime: datetime,
		Currency: models.CurrencyBTC,
		Price:    closePrice,
	}, nil
}
