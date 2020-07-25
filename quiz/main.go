package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	Qstn, Ans string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV File : %s\n", *csvFileName))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the csv")
	}
	correct := 0

	problems := parseLines(lines)
	// timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) //time for all question
	for i, p := range problems {
		fmt.Printf("Problem #%d : %s = \n", i+1, p.Qstn)
		timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) //time for each question
		AnswerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s \n", &answer)
			AnswerCh <- answer
		}()
		select {
		case <-timer.C:

		case answer := <-AnswerCh:

			if answer == p.Ans {
				correct++
			}
		}

	}
	fmt.Printf("You Scored %d out of %d\n", correct, len(problems))

}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			Qstn: line[0],
			Ans:  line[1],
		}
	}
	return ret
}
