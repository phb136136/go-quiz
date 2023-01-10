package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Question struct {
	Ask    string
	Answer string
}

func storeQuestions(data [][]string) []Question {
	var questions []Question
	for i, line := range data {
		if i > 0 { // omit header line

			// Use a string literal
			qs := Question{
				Ask:    line[0],
				Answer: line[1],
			}
			questions = append(questions, qs)
		}
	}
	return questions
}

func runQuestions(questions []Question) {
	for _, qs := range questions {
		var userInput string
		fmt.Printf("%s\n", qs.Ask)
		// https://gosamples.dev/read-user-input/
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			log.Fatal(err)
		}
		if userInput != qs.Answer {
			fmt.Printf("User gave the wrong answer\n")
			fmt.Printf("Expected: %s, Got: %s\n", qs.Answer, userInput)
			break
		}

	}
}

func main() {
	// open file
	filename := "problems.csv"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	questions := storeQuestions(data)

	// print the array testing this works
	// fmt.Printf("%+v\n", questions)
	runQuestions(questions)
}
