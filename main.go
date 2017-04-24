package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	historyFile := flag.String("history", "history.csv", "history file location")

	flag.Parse()
	csvFile, err := os.Open(*historyFile)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	reader := csv.NewReader(csvFile)
	d, err := getFirstLine(reader)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Println(d)
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