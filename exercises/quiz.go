package exercises

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Quiz func
// -csv flag to specify csv file (default file is extras/problems.csv)
// -limit flag to specify time limit in seconds
func Quiz() {
	// -csv flag
	csvFilename := flag.String("csv", "extras/problems.csv", "a csv file in the format of 'question, answer'")
	// -limit flag
	timeLimit := flag.Int("limit", 30, "time limit for quiz in seconds")
	flag.Parse()
	// opening the csv file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)
	// fmt.Println(problems)

	// time.Duration is a type alias of int
	// sends current time in its channel "timer.C"
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	counter := 0
problemLoop:
	for i, p := range problems {
		fmt.Printf("problem #%d: %s = ", i+1, p.question)

		// answer channel to get answer entered by the user
		answerCh := make(chan string)

		// creating a anonymous goroutine function that takes
		// input from user and sends input through answer channel
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			// send user answer in a channel
			answerCh <- answer
		}()
		select {
		// if time limit specified in the flag is finished then loop is ended
		case <-timer.C:
			fmt.Printf("\n")
			break problemLoop
		// if answer entered by the user is entered then counter is increased by 1
		case answer := <-answerCh:
			if answer == p.answer {
				counter++
			}
		}
	}
	fmt.Printf("You scored %d out of %d. \n", counter, len(problems))
}

type problem struct {
	question string
	answer   string
}

// returns a slice of struct for QnA from 2D slice of QnA as argument
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// exits by using system exit status 1 and printing error message
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
