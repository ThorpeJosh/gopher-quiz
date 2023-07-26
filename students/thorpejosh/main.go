package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the quiz!\n Press q at anytime to quit")

	csvFile, err := os.Open("../../problems.csv")
	if err != nil {
		log.Fatalf("Error reading CSV file: %s", err)
	}
	// remember to close the file at the end of the program
	defer csvFile.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(csvFile)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// data is a 2D string slice that mirrors the csv file.
	for idx, line := range data {
		question := strings.TrimSpace(line[0])
		answer := strings.TrimSpace(line[1])
		fmt.Printf("Question %d: What is %s ?\n", idx, question)

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		if input == "q" {
			os.Exit(0)
		}

		if input == answer {
			fmt.Printf("Correct\n")
		} else {
			fmt.Printf("Incorrect\n")
		}
	}
}
