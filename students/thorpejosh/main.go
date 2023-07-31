package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvPath := "../../problems.csv"
	timeout := time.Duration(10) // Timeout in seconds

	// data is a 2D string slice that mirrors the csv file.
	data := readCSV(csvPath)
	totalQuestions := 0
	correctAnswers := 0

	fmt.Println("Welcome to the quiz!\n Press q at anytime to quit")
	for idx, line := range data {

		question := strings.TrimSpace(line[0])
		answer := strings.TrimSpace(line[1])
		fmt.Printf("Question %d: What is %s ?\n", idx, question)

		inputChannel := make(chan string)
		timeoutChannel := make(chan int)

		go func() {
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			inputChannel <- strings.TrimSpace(input)
		}()

		go func() {
			time.Sleep(timeout * time.Second)
			timeoutChannel <- 1
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
		case <-timeoutChannel:
			fmt.Printf("You took too long!")
			os.Exit(0)
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
