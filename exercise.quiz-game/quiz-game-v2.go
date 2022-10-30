// updated version of my personal answer now that I've seen the proper solutions

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// * the main difference between my solution and the officia
	// * are that I used a goroutine for my timer instead of questions
	// initialize cli flags
	csvFileName := flag.String("csv", "problems.csv", "uses relative or absolute file path to read a file with problems in the format of 'question, answer'.")
	timeLimit := flag.Int("time", 30, "set the time limit in seconds.")
	shuffle := flag.Bool("shuffle", false, "shuffle the order of problems every time the quiz is run.")
	flag.Parse()

	// initialize variables
	lines := readFile(*csvFileName)
	if *shuffle {
		lines = shuffleLines(lines)
	}
	// * and that I didn't know about structs lol
	problemList := parseLines(lines)
	correctAnswerCounter := 0

	// ask user to confirm before starting quiz
	fmt.Println("press <enter> to start timer and begin quiz")
	fmt.Scanln() // wait for <enter> key
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemLoop:
	for questionCount, problemStruct := range problemList {
		fmt.Printf("problem #%d: %s \n", questionCount+1, problemStruct.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			// * and that I didn't know about the Scanf() method
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("It's been %vs and you've run out time.", *timeLimit)
			break problemLoop
		case answer := <-answerChannel:
			if answer == problemStruct.answer {
				correctAnswerCounter++
			}
		}
	}
	fmt.Printf("you managed to get %v out of %v questions correct, congratulations.\n", correctAnswerCounter, len(problemList))
}

func readFile(csvFileName string) [][]string {
	// open a csv file
	file, err := os.Open(csvFileName)
	if err != nil {
		fmt.Printf("failed to open csv file %q\n", csvFileName)
		panic(err)
	}

	// defer closing file till function is done
	defer file.Close()

	// parse a csv file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("failed to parse provided csv file %q\n", csvFileName)
	}

	// export output
	return lines
}

func shuffleLines(lines [][]string) [][]string {
	// for every line, swap the index of two lines based on unix time
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
	return lines
}

func parseLines(lines [][]string) []problem {
	// parse 2dslice output into 1dslice and struct
	parsed_lines := make([]problem, len(lines))
	for line_num, line := range lines {
		parsed_lines[line_num] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return parsed_lines
}

type problem struct {
	question string
	answer   string
}
