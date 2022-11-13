// time.AfterFunc() Function in Golang With Examples - GeeksforGeeks
// https://www.geeksforgeeks.org/time-afterfunc-function-in-golang-with-examples/

package main

import (
	"fmt"
	"time"
)

func main() {
	// this is a more advanced but still generic timer that makes use of go routines
	// and does in fact appear to non-blocking, next step would be to run some demo code
	// from my exercise through it, to see if it'll actually work the way i think it will

	// define a go-channel with method make()
	channel := make(chan int)

	// define method AfterFunc() with it's parameters
	time.AfterFunc(6*time.Second, func() {
		// print out statement to confirm func was called
		fmt.Println("func called after 6 seconds.")
		// https://gobyexample.com/channels
		// input the following into channel
		channel <- 30
	})

	// pretty sure this is non-blocking :O
	for {
		select {
		case n := <-channel:
			fmt.Println("output from channel arriving: ", n)
			fmt.Println("done!")
			// exit program after channel returns output
			return
		// return the following by default
		default:
			fmt.Println("time to wait 2s")
			time.Sleep(2 * time.Second)
		}
	}
}

/*
Output:
	time to wait 2s
	time to wait 2s
	time to wait 2s
	func called after 6 seconds
	output from channel arriving
	done!
*/
