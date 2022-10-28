// time.AfterFunc() Function in Golang With Examples - GeeksforGeeks
// https://www.geeksforgeeks.org/time-afterfunc-function-in-golang-with-examples/

package main

import (
	"fmt"
	"time"
)

func main() {
	// this is just a generic basic timer, to make sure the simple parts actually work
	// in this example, the output is returned after 3s and then the returned timer
	// to cancel the call to function timer() using method Stop(), after that
	// the program sleeps for 10s and exits the program

	// define method AfterFunc()'s wait_for parameter
	wait_for := time.Duration(3) * time.Second

	// define method AfterFunc()'s do_action parameter
	do_action := func() {
		// print out statement to confirm it was called
		fmt.Printf("function called by method AfterFunc() after %s \n", wait_for)
	}

	// call method AfterFunc() with it's pre-defined parameters
	timer := time.AfterFunc(wait_for, do_action)

	// use returned timer to cancel the call to function timer()
	defer timer.Stop()

	// call method sleep() to delay exiting the program
	time.Sleep(10 * time.Second)
}
