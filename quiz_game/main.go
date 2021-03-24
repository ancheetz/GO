package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	getCsvProblems()

}

func getCsvProblems() {

	csvfile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Could't open the file!", err)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %v Answer: %v\n", record[0], record[1])
	}

}
