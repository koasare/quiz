package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Type Definitions

type Problems struct {
	/* Problems Struct to Sort CSV file in question and answer */
	question string
	answer   string
}

// Functions

func parseLines(records [][]string) []Problems {
	/* func parseLines takes string array from csv file and parses
	each line to create a seperation between questions and answers
	*/
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
	/*Exit function to condition
	exits in the event of an error*/
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
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
	answerTimer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct_answer := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s =  \n", i+1, problem.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-answerTimer.C:
			fmt.Printf("You scored %d out of %d questions", correct_answer, len(problems))
			return
		case answer := <-answerChannel:
			if answer == problem.answer {
				correct_answer++
			}
		}
	}
	fmt.Printf("You scored %d out of %d questions", correct_answer, len(problems))
}
