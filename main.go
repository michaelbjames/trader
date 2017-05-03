package main

import (
	"encoding/csv"
	"fmt"
	"local/trader/analysis"
	"local/trader/mappers"
	"local/trader/models"
	"local/trader/parsers"
	"os"
	"strconv"

	"github.com/pkg/errors"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "historical",
			Usage: "run an analysis on historical data",
			Action: func(c *cli.Context) error {
				principalStr := c.Args().First()
				brokerageName := c.Args().Get(1)
				historyFile := c.Args().Get(2)
				return historical(principalStr, brokerageName, historyFile)
			},
		},
	}

	app.Run(os.Args)
}

func historical(principalStr string, brokerageStr string, historyFileName string) error {
	principal, err := strconv.ParseFloat(principalStr, 64)
	if err != nil {
		return errors.Wrap(err, "principal is not a number")
	}
	broker, err := mappers.ParseBroker(brokerageStr)
	if err != nil {
		return errors.Wrap(err, "could not parse broker")
	}
	csvFile, err := os.Open(historyFileName)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.Wrap(err, "Unable to open file")
	}
	reader := csv.NewReader(csvFile)
	closes, err := parsers.CoindeskMarketClose(reader)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	_ = newHistoricalHistory(principal, broker, closes)

	analysis.TwoDayStreaks(closes)
	return nil
}

func newHistoricalHistory(principal float64, broker models.BrokerageName, events []models.Event) models.History {
	switch broker {
	case models.BrokerAnalysis:
		_ = models.AnalysisInit{
			Source:    events,
			Principal: principal,
		}
	}
	return nil
}
