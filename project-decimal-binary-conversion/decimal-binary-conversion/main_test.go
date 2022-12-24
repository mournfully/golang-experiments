package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestConvertParsedOut(t *testing.T) {
	readOut, err := readFile("input.csv")
	if err != nil {
		fmt.Printf("main: %v", err)
		return
	}

	shuffleOut := shuffleLines(readOut)
	parsedOut := parseLines(shuffleOut)
	// to see output use, "go test ./main_test.go -v"
	// t.Logf("%# v\n", pretty.Formatter(parsedOut))
	if 
}

func readFile(csvFileName string) ([][]string, error) {
	file, err := os.Open(csvFileName)
	if err != nil {
		return nil, fmt.Errorf("readFile: failed to open csv file: %w", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("readFile: failed to parse csv file: %w", err)
	}

	return lines, nil
}

func shuffleLines(lines [][]string) [][]string {
	// for every line, swap the index of two lines based on unix time
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
	return lines
}

func parseLines(input [][]string) []number {
	parsed_lines := make([]number, len(input))
	for line_num, line := range input {
		parsed_lines[line_num] = number{
			decimal: line[0],
			binary:  line[1],
		}
	}
	return parsed_lines
}

type number struct {
	decimal string
	binary  string
}
