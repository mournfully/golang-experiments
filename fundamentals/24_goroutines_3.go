// this example, is stupid...
// ima delete this later
// holy shit, i just ate my own words lol - this is important

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d starting \n", id)
	// use sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("worker %d finished \n", id)
}

func main() {
	// https://gobyexample.com/waitgroups
	// wait for all goroutines launched from here to finish
	// explicitly pass waitgroup into functions thru pointers
	var wg sync.WaitGroup
	// launch 5 goroutines
	for i := 1; i <= 5; i++ {
		// increment kwaitgroup counter for each goroutine
		wg.Add(1)
		// this looks odd and i'm not sure why it works but its supposed to...
		// create a new 'i' each iteration to ensure no reusing of same value in each goroutine
		i := i
		// wrap worker() in closure that notifies waitgroup on completion
		// to avoid worker having to be aware of concurrency complexity
		// this is certainly not perfect, as there's no easy way to propogate worker errorss
		go func() {
			defer wg.Done()
			// https://gobyexample.com/goroutines
			worker(i)
		}()
	}
	// wait for all goroutines to complete before exiting
	wg.Wait()

	/*
		Output:
		 worker 5 starting
		 worker 2 starting
		 worker 3 starting
		 worker 4 starting
		 worker 1 starting
		 worker 1 finished
		 worker 2 finished
		 worker 5 finished
		 worker 4 finished
		 worker 3 finished
	*/
}
