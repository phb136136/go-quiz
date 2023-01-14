package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Ask    string
	Answer string
}

// Parses the csv and stores the data
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

// Randomises the questions
func randomiseQuestions(questions []Question) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

}

func runQuestions(questions []Question) {
	var numCorrect, numWrong int = 0, 0
	randomiseQuestions(questions)
	for _, qs := range questions {
		var userInput string
		fmt.Printf("%s\n", qs.Ask)
		// https://gosamples.dev/read-user-input/
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(strings.TrimSpace(userInput)) == qs.Answer {
			numCorrect++
		} else {
			numWrong++
			// fmt.Printf("User gave the wrong answer\n")
			// fmt.Printf("Expected: %s, Got: %s\n", qs.Answer, userInput)
			// break
		}

	}
	fmt.Printf("You got %d answers correct out of a total of %d questions\n", numCorrect, numWrong+numCorrect)
}

func main() {
	// open file
	filename := "problems.csv"
	timeLimitSeconds := 15

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
	if len(os.Args) > 1 {
		timeout, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		timeLimitSeconds = timeout
	}

	var waitStart string
	fmt.Println("Press any key to start")
	_, error := fmt.Scan(&waitStart)
	if error != nil {
		log.Fatal(error)
	}

	// Using go channels
	c2 := make(chan string, 1)
	go func() {
		runQuestions(questions)
		c2 <- "result 2"
	}()

	select {
	case <-c2:
		fmt.Println("FINISHED: YOU WIN")
	case <-time.After(time.Duration(timeLimitSeconds) * time.Second):
		fmt.Println("timeout: you LOST")
	}
}
