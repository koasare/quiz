package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	fmt.Println(records)

}
