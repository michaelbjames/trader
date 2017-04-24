package parsers

import (
	"encoding/csv"
	"local/trader/models"
)

const coindeskDateFmt string = "2006-01-02 15:04:05"

func CoindeskMarketClose(reader *csv.Reader) ([]models.MarketClose, error) {
	options := options{
		DateFmt:  coindeskDateFmt,
		Currency: models.CurrencyBTC,
	}
	return datevalueParse(reader, options)
}
