package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	csvFile := flag.String("csv", "problems.csv", "a file containing records as 'questions, answers'")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Sprintf("File would not open, %s\n", *csvFile)
		os.Exit(1)
	}
	r := csv.NewReader(file)
}
