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
	- [ ] ask user to press enter to start a 30 second quiz by default and add [flag -t] to change time limit
			how does calling timer() from main() and having it change a global flag after n seconds
			that would in turn flip an if-else in the for-loop?
			although, that might not kick the user if they're mid-question
			OMG, GO ROUTINES AND CHANNELS!!!
	- [ ] stop quiz immediately even if mid-question as soon as time limit is reached

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
	"strings"
)

// helper for streamlining error checks during dev
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// used top answer for parsing data - https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
// understanding output - https://www.dotnetperls.com/csv-go
func read(fp string) [][]string {
	// fp = file path, f = file, err = error, rec = record
	// load a csv file
	f, err := os.Open(fp)
	if err != nil {
		err := fmt.Errorf("unable to read csv file %q\n %v", fp, err)
		fmt.Println(err.Error())
	}
	// defer till surrounding function returns output
	defer f.Close()

	// create a new reader and start parsing
	// readAll() outputs a 2d slice of slices
	csv := csv.NewReader(f)
	rec, err := csv.ReadAll()
	if err != nil {
		err := fmt.Errorf("unable to parse csv file %q\n %v", fp, err)
		fmt.Println(err.Error())
	}

	return rec
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

func timer() {

}

// program entry point - must be main
func main() {
	counter := 0
	// fp = file path, csv_out = csv output
	fp := flag.String("fp", "problems.csv", "Uses relative or absolute file path to select a file with problems for the quiz read.")
	flag.Parse()
	// outputs 2d slice (a slice of slices)
	csv_out := read(*fp)
	//// fmt.Printf("%v lines | %s \n", len(output), output)

	// https://gobyexample.com/slices
	// https://www.dotnetperls.com/csv-go
	// ln = line number, l = line, qn = question, ans = answer, in = input
	for line_num := range csv_out {
		line := csv_out[line_num] // [[5+5 10] [1+2 2]] -> [5+5 10]
		question := line[0]       // [5+5 10] -> 5+5
		ans := line[1]            // [5+5 10] -> 10

		input := ask(question)
		if ans == input {
			//// fmt.Printf("expected: %s | input: %s | result: true \n", answer, input)
			counter++
		}
		//// else {
		//// 	// testing
		//// 	fmt.Printf("expected: %s | input: %s | result: false \n", answer, input)
		//// }
	}
	fmt.Printf("You've reached the end of the quiz, you got %v out of %v questions correct", counter, len(csv_out))
}
