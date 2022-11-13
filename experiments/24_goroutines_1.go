// this certainly works, but i'm not sure if it'd fit into my current projects code
// without extensive refactoring

package main

import (
	"fmt"
	"sync"
)

// sample function for learning
func f(wg *sync.WaitGroup, from string) {
	// decrease waitgroup counter by 1
	wg.Done()
	for i := 0; i <= 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func execute() {
	// waitgroup exports 3 methods:
	// - Add(int), increase waitgroup counter by int
	// - Done(), decreases waitgroup counter by 1
	// - Wait(), blocks execution till counter = 0

	// define waitgroup to var wg
	wg := new(sync.WaitGroup)
	// increase waitgroup counter by 2 because we're running 2 goroutines
	wg.Add(2)

	go f(wg, "direct")    // default func call
	go f(wg, "goroutine") // invoked as a concurrent goroutine

	wg.Wait()
}

func main() {
	// https://stackoverflow.com/questions/16228887/why-does-fmt-println-inside-a-goroutine-not-print-a-line
	// https://www.geeksforgeeks.org/using-waitgroup-in-golang/
	// https://gobyexample.com/channel-directions
	/*
		// if you ran go routtines like this there would be no output because
		// as soon as you launch both goroutines your main function just got terminated
		go f("direct")    // default func call
		go f("goroutine") // invoked as a concurrent goroutine
	*/

	// launch runners
	execute()

	/*
		Output
		 goroutine :  0
		 goroutine :  1
		 goroutine :  2
		 goroutine :  3
		 direct :  0
		 direct :  1
		 direct :  2
		 direct :  3
	*/
}
