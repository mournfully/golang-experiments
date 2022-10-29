/**
- References
	https://gophercises.com/#signup
	https://github.com/gophercises/quiz
	https://gobyexample.com/

- Part 1 (Basic)
	- [x] read data from problems.csv by default and add [flag -fp] to change input file
	- [x] manipulate data into seperate questions and associated answers
	- [x] ask user questions and track only if they were correct
		- if answer == correct --> counter++ ...then next question
		- if answer != correct --> ...next question
		- if answer == nil --> ...next question
	- [x] at the end output score of # of correct/# of questions

- Part 2 (Advanced)
	- [x] ask user to press enter to start a 30 second quiz by default and add [flag -t] to change time limit
	- [x] stop quiz immediately even if mid-question as soon as time limit is reached
		how does calling timer() from main() and having it change a global flag after n seconds
		that would in turn flip an if-else in the for-loop?
		although, that might not kick the user if they're mid-question
		OMG, GO ROUTINES AND CHANNELS!!!

- Part 3 (Bonus)
	- [ ] sanatize user inputs with 'strings' package
	- add [flag -s] to shuffle questions around every run
	- use golang library 'cobra' to display real-time countdown while quiz is running
	- create unit tests for exercise
*/

// package name can be an arbitrary name unless it's an entrypoint for your program
// https://forum.golangbridge.org/t/error-package-command-line-arguments-is-not-a-main-package/25851/4
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// used top answer for parsing data - https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
// understanding output - https://www.dotnetperls.com/csv-go
func read(file_path string) [][]string {
	// fp = file path, f = file, err = error, res = result
	// load a csv file
	file, err := os.Open(file_path)
	if err != nil {
		err := fmt.Errorf("unable to read csv file %q\n %v", file_path, err)
		fmt.Println(err.Error())
	}
	// defer till surrounding function returns output
	defer file.Close()

	// create a new reader and start parsing
	// readAll() outputs a 2d slice of slices
	csv := csv.NewReader(file)
	result, err := csv.ReadAll()
	if err != nil {
		err := fmt.Errorf("unable to parse csv file %q\n %v", file_path, err)
		fmt.Println(err.Error())
	}

	return result
}

// https://dev.to/tidalmigrations/interactive-cli-prompts-in-go-3bj9
// https://freshman.tech/snippets/go/read-console-input/
func ask(question string) string {
	fmt.Printf("%s ", question)
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		err := fmt.Errorf("an error occured while reading input %v", err)
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(answer)
}

// channel chan<- string,
func timer(wg *sync.WaitGroup, duration int, correct_answer_counter chan string, total_num_of_question int) {
	// t = time, channel = # of correct answer, total = # of total questions,
	// convert flag -t's int output to something time.AfterFunc() can understand and define method AfterFunc() with it's input parameter set
	time.AfterFunc(time.Duration(duration)*time.Second, func() {
		// print out statement to confirm func was called
		fmt.Printf("\nfunc called after %vs \n", duration)
		// this is kind of yucky but... runs the following when user doesn't finish before timer ends
		fmt.Printf("It's been %vs and you've run out of time, you managed to get %v out of %v questions correct \n", duration, <-correct_answer_counter, total_num_of_question)
		wg.Done()
		os.Exit(0)
	})
}

// program entry point - must be main
func main() {
	// initialize cli flags
	// fp = file path, csv_out = csv output
	fp := flag.String("fp", "problems.csv", "Uses relative or absolute file path to select a file with problems for the quiz read.")
	t := flag.Int("t", 30, "Sets the timer's duration in seconds.")
	flag.Parse()
	// create waitgroup for goroutine timer()
	wg := new(sync.WaitGroup)
	wg.Add(1)
	// create buffered channel because... i forgot :skull:
	channel := make(chan string, 1)
	channel <- "0"

	// call functions
	ask("press <enter> to start timer and begin quiz")
	counter := 0
	csv_out := read(*fp) // outputs 2d slice (a slice of slices)
	//// fmt.Printf("%v lines | %s \n", len(output), output)
	go timer(wg, *t, channel, len(csv_out))

	// https://gobyexample.com/slices
	// https://www.dotnetperls.com/csv-go
	// ln = line number, l = line, qn = question, ans = answer, in = input
	for line_num := range csv_out {
		// parsing csv to usable data w/ examples on what they do
		line := csv_out[line_num] // [[5+5 10] [1+2 2]] -> [5+5 10]
		question := line[0]       // [5+5 10] -> 5+5
		answer := line[1]         // [5+5 10] -> 10

		// simple question validation
		input := ask(question)
		if answer == input {
			counter++
			fmt.Printf("expected: %s | input: %s | result: correct \n", answer, input)
		} else {
			fmt.Printf("expected: %s | input: %s | result: incorrect \n", answer, input)
		}

		// run below `counter++` and ask() so that when user's mid-question and timer() uses os.Exit()
		// counter remains same as before whichever question user was in the middle of
		// explicitly empty channel before sending new data
		<-channel
		channel <- strconv.Itoa(counter)
	}
	// runs when user gets through every question before timer ends
	fmt.Printf("You've reached the end of the quiz, you got %v out of %v questions correct \n", counter, len(csv_out))

	// if im going to use wg.Wait(), what happens if user finishes quiz before 30s are up?
	wg.Wait()
}
