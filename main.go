package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type quizItem struct {
	q string
	a string
}

func parseLine(csvData [][]string) []quizItem {
	result := make([]quizItem, len(csvData))
	for i, line := range csvData {
		result[i] = quizItem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return result
}

func main() {

	// SECTION: setting up the flags
	var csvFileName string
	var timeLimit int
	var shuffle bool

	flag.StringVar(&csvFileName, "csv", "problems.csv", "Enter the csv that is in the format: 'question, answer'")
	flag.IntVar(&timeLimit, "timelimit", 30, "Configure the time limit for each question.")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle the questions in the quiz.")
	flag.Parse()

	// SECTION: Reading and Processing the file data
	file, err := os.Open(csvFileName)
	if err != nil {
		fmt.Print("Could not parse the file name.")
		os.Exit(1)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Could not read the given file")
		os.Exit(1)
	}

	problemList := parseLine(lines)

	// SECTION: Shuffling of the Deck
	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problemList), func(i, j int) {
			problemList[i], problemList[j] = problemList[j], problemList[i]
		})
	}

	// SECTION: Rendering the quiz
	var score float32 = 0
	input := make(chan string)
	reader := bufio.NewReader(os.Stdin)

	for i, problem := range problemList {

		fmt.Printf("Proble #%d: %s = ", i+1, problem.q)

		go func() {
			answer, _ := reader.ReadString('\n')
			input <- strings.TrimSpace(answer)
		}()

		select {
		case result := <-input:
			if problem.a == result {
				score++
			} else {
				fmt.Println("Worng")
				score = score - 0.5
			}
		case <-time.After(time.Duration(timeLimit) * time.Second):
			fmt.Println("Time is up. Moving On!!")
		}
	}
	fmt.Printf("Score: %.2f", score)
}
