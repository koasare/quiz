package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Type Definitions
type Problems struct {
	question string
	answer   string
}

// Functions
func parseLines(records [][]string) []Problems {
	result := make([]Problems, len(records))
	for i, record := range records {
		result[i] = Problems{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
	}
	return result
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println(err)
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFileName))
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		exit("Failed to Parse File")
	}
	problems := parseLines(records)
	//fmt.Println(problems)
	correct_answer := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s =  \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct_answer++
		} else {
			fmt.Println("Oops, wrong Answer!")
			fmt.Println("The correct answer is:", problem.answer)
		}
	}
	fmt.Printf("You scored %d out of %d questions", correct_answer, len(problems))

}
