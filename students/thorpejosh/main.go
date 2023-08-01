package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	csvPath := flag.String("csv", "../../problems.csv", "Path to the csv file that contains the questions,answers")
	timeout := flag.Int("timeout", 5, "Timout for each question in seconds") // Timeout in seconds
	flag.Parse()

	// data is a 2D string slice that mirrors the csv file.
	data := readCSV(*csvPath)
	totalQuestions := 0
	correctAnswers := 0

	fmt.Println("Welcome to the quiz!\n Press q at anytime to quit")
	for idx, line := range data {

		question := strings.TrimSpace(line[0])
		answer := strings.TrimSpace(line[1])
		fmt.Printf("Question %d: What is %s ?\n", idx, question)

		inputChannel := make(chan string)

		go func() {
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			inputChannel <- strings.TrimSpace(input)
		}()

		select {

		case input := <-inputChannel:
			if input == "q" {
				os.Exit(0)
			}

			if input == answer {
				fmt.Printf("Correct\n")
				correctAnswers++
			} else {
				fmt.Printf("Incorrect\n")
			}
			totalQuestions++
			os.Exit(0)
		case <-time.After(time.Duration(*timeout) * time.Second):
			fmt.Println("You took too long!")
		}
	}
	fmt.Printf("Well done! You achieved a score of %d/%d", correctAnswers, totalQuestions)
}

func readCSV(path string) [][]string {
	csvFile, err := os.Open(path)
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
	return data
}
