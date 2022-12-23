package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	a := 0

	// Scan the given file line by line
	lineScanner := bufio.NewScanner(file)
	for lineScanner.Scan() {
		// Scan the given line word by word
		wordScanner := bufio.NewScanner()
		wordScanner.Split(bufio.ScanWords)

		fmt.Println(wordScanner.Text())
		fmt.Println(wordScanner.Text())

		fmt.Println(lineScanner.Text())

		a++
		if a > 10 {
			break
		}
	}
}
