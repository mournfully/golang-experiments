package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words, err := scanWords("input.txt")
	if err != nil {
		panic(err)
	}

	for _, word := range words {
		fmt.Println(word)
	}
}

func scanWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}
