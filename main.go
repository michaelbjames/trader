package main

import (
	"encoding/csv"
	"fmt"
	"local/trader/gateways"
	"local/trader/mappers"
	"local/trader/models"
	"local/trader/parsers"
	"local/trader/updaters"
	"math/big"
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
	brokerName, err := mappers.ParseBroker(brokerageStr)
	if err != nil {
		return errors.Wrap(err, "could not parse broker")
	}
	csvFile, err := os.Open(historyFileName)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.Wrap(err, "Unable to open file")
	}
	reader := csv.NewReader(csvFile)
	events, err := parsers.CoindeskMarketClose(reader)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	analysisInit := models.AnalysisInit{
		Source:    events,
		Principal: big.NewFloat(principal),
	}
	broker := gateways.New(brokerName, analysisInit)
	historyCore := newHistoricalHistory(big.NewFloat(principal), brokerName, events)
	for _, day_close := range events {
		historyCore, err := historyCore.SaveEvent(day_close)
		if err != nil {
			fmt.Printf("Save Event failed: %v", err)
			panic("could not save event")
		}
		action, err := historyCore.TakeAction()
		if err != nil {
			fmt.Printf("Take Action failed: %v", err)
			panic("could not take action")
		}
		broker.Trade(action)
	}

	// analysis.TwoDayStreaks(closes)
	return nil
}

func newHistoricalHistory(principal *big.Float, broker models.BrokerageName, events []models.Event) models.History {
	switch broker {
	case models.BrokerAnalysis:

		engineState := updaters.NewManicMomentum(principal)
		return engineState
	}
	return nil
}
