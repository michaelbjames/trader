package main

import (
	"encoding/csv"
	"fmt"
	"local/trader/analysis"
	"local/trader/parsers"
	"os"
	"time"

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
				fmt.Println("added task: ", c.Args().First())
				return historical(principalStr, brokerageName, historyFile)
			},
		},
	}

	app.Run(os.Args)
}

func historical(principalStr string, brokerageStr string, historyFileName string) error {
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
	analysis.TwoDayStreaks(closes)
	return nil
}

func getFirstLine(reader *csv.Reader) ([]string, error) {
	l1, err := reader.Read()
	if err != nil {
		return nil, err
	}
	_, err = time.Parse("2006-01-02 15:04:05", l1[0])
	if err == nil {
		return l1, nil
	}
	return getFirstLine(reader)
}
