/**
* https://gophercises.com/#signup
* https://github.com/gophercises/quiz
* https://gobyexample.com/

* Plugins
	* encodig/csv

* Plan
	* NOTE: CSV files may have questions with commas in them.
	* Eg: "what 2+2, sir?",4 is a valid row in a CSV.
	* I suggest you look into the CSV package in Go and don't try to write your own CSV parser.
		* * finished importing "encoding/csv" package

	* The CSV file should default to problems.csv
	* But the user should be able to customize the filename via a flag.
		* *  oh okay so just a "golang command-line flag w/ default of problems.csv"
	* Create a program that will read in a quiz provided via a CSV file
		* * done, got CSV Parsing with cli flags implemented

	* And will then give the quiz to a user
		* TODO: we'll need to take csv output and manipulate it somehow
	* Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.
		* * this sounds like a dumb "for loop"
		* TODO: hmm okay i'll get to it
	* You can assume that quizzes will be relatively short (< 100 questions)
		* * ok so definitely a "for loop"
		* * and we don't need optimizations like "reading csv line by line on demand" with an input so small

	* And will have single word/number answers.
		* * not sure what to do with this
		* ! m := make(map[string]int)
		* ! this will be a problem ^ so we probably cant use "maps"
		* TODO: wait how do we prompt? with promptui library?

	* While, keeping track of how many questions they get right and how many they get incorrect.
	* At the end of the quiz the program should output the total number of questions
	* correct and how many questions there were in total.
	* Questions given invalid answers are considered incorrect.


	Part 1:
	format: problem, answer
	read problems.csv file by default - let user customize filename via flag

	ask question and prompt for an answer
	start next answer immediatley after

	keep track of how many questions were correct or incorrect
	at the end of the quiz output: correct answers/total questions
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
func read_csv(fp string) [][]string {
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

// program entry point - must be main
func main() {
	counter := 0
	// fp = file path, out = output
	fp := flag.String("i", "problems.csv", "Selects a file from which the problems for the quiz are read. A different file can be set with it's relative or absolute path")

	flag.Parse()
	// outputs 2d slice (a slice of slices)
	out := read_csv(*fp)
	//// fmt.Printf("%v lines | %s \n", len(output), output)

	// https://gobyexample.com/slices
	// https://www.dotnetperls.com/csv-go
	// ln = line number, l = line, qn = question, ans = answer, in = input
	for ln := range out {
		line := out[ln]     // [[5+5 10] [1+2 2]] -> [5+5 10]
		question := line[0] // [5+5 10] -> 5+5
		ans := line[1]      // [5+5 10] -> 10

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
	fmt.Printf("You've reached the end of the quiz, you got %v out of %v questions correct", counter, len(out))
}
