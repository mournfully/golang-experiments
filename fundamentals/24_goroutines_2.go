// wtf why did v1 work but not this?
// oh, figured it out.
// wg.Done() should be at the end of f()
// not the start lol

package main

import (
	"fmt"
	"sync"
	"time"
)

// sample function for learning
func f(wg *sync.WaitGroup, from string) {
	// decrease waitgroup counter by 1
	for i := 0; i <= 3; i++ {
		fmt.Println(from, ": ", i)
		// use sleep to simulate an expensive task
		time.Sleep(time.Second)
	}
	// must be at end of runner function lol
	wg.Done()
}

func main() {
	// https://stackoverflow.com/questions/16228887/why-does-fmt-println-inside-a-goroutine-not-print-a-line
	// https://www.geeksforgeeks.org/using-waitgroup-in-golang/
	// https://gobyexample.com/channel-directions

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

	/*
		Output:
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
