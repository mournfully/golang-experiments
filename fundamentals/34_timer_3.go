// time.AfterFunc() Function in Golang With Examples - GeeksforGeeks
// https://www.geeksforgeeks.org/time-afterfunc-function-in-golang-with-examples/

// Go by Example: Non-Blocking Channel Operations
// https://gobyexample.com/non-blocking-channel-operations

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

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
func timer(wg *sync.WaitGroup, t int) {
	// convert flag -t's int output to something time.AfterFunc() can understand
	// define method AfterFunc() with it's input parameter set
	time.AfterFunc(time.Duration(t)*time.Second, func() {
		// print out statement to confirm func was called
		fmt.Printf("func called after %vs \n", t)
		// channel <- "idk"
		wg.Done()
	})
}

func main() {
	// initialize new variables for testing
	// channel := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// initialize variables from past commits
	t := flag.Int("t", 2, "Time, in seconds")
	csv_out := [][]string{
		{"5+5", "10"},
		{"1+1", "2"},
	}
	flag.Parse()
	counter := 0
	fmt.Println(csv_out)

	go timer(wg, *t)

	for line_num := range csv_out {
		line := csv_out[line_num] // [[5+5 10] [1+2 2]] -> [5+5 10]
		question := line[0]       // [5+5 10] -> 5+5
		answer := line[1]         // [5+5 10] -> 10

		input := ask(question)
		// go ask(wg, question, answer)
		// return via channel result of correct or not?
		if answer == input {
			counter++
			fmt.Printf("expected: %s | input: %s | result: correct \n", answer, input)
		} else {
			fmt.Printf("expected: %s | input: %s | result: incorrect \n", answer, input)
		}
	}
	fmt.Printf("You've reached the end of the quiz, you got %v out of %v questions correct \n", counter, len(csv_out))

	// if im going to use wg.Wait(), what happens if user finishes quiz before 30s are up?
	wg.Wait()
}
