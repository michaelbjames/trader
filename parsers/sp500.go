package parsers

import (
	"encoding/csv"
	"local/trader/models"
)

const sp500DateFmt string = "2006-01-02"

func SP500(reader *csv.Reader) ([]models.MarketClose, error) {
	options := options{
		DateFmt:  sp500DateFmt,
		Currency: models.CurrencyAbstract,
	}
	return datevalueParse(reader, options)
}
