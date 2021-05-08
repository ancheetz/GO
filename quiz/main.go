package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	csvFile := flag.String("csv", "problems.csv", "the csv file in the format of question,answer")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to print csv file: %s\n", *csvFile))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == p.a {
			correct++
			fmt.Println("Correct!")
		}
	}
	fmt.Printf("You scored %d out of %d, in this quiz!\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
