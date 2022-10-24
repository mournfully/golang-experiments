/*
Plugins
	encodig/csv
	os - read file from os
	fmt - idk, why not
	math - not sure

Part 1:
	format: problem, answer
	read problems.csv file by default - let user customize filename via flag

	ask question and prompt for an answer
	start next answer immediatley after

	keep track of how many questions were correct or incorrect
	at the end of the quiz output: correct answers/total questions
*/

// https://forum.golangbridge.org/t/error-package-command-line-arguments-is-not-a-main-package/25851/4
// package name can be an arbitrary name unless it's an entrypoint
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// helper for streamlining error checks during dev
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
func read_csv(fp string) [][]string {
	// fp = file path, f = file, err = error, rec = record
	// open csv file
	f, err := os.Open(fp)
	if err != nil {
		err := fmt.Errorf("unable to read csv file %q\n %v", fp, err)
		fmt.Println(err.Error())
	}
	// defer till surrounding function returns output
	defer f.Close()

	// read csv values line by line
	csv := csv.NewReader(f)
	rec, err := csv.ReadAll()
	if err != nil {
		err := fmt.Errorf("unable to parse csv file %q\n %v", fp, err)
		fmt.Println(err.Error())
	}

	// output records
	return rec
}

// program entry point - must be main
func main() {
	// fp = file path, out = output
	fp := flag.String("i", "problems.csv", "Selects a file from which the problems for the quiz are read. A different file can be set with it's relative or absolute path")
	flag.Parse()
	out := read_csv(*fp)
	fmt.Println(out)
}
