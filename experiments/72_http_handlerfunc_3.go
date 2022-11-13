/*
	Most of the time using a function as a handler like this works well. But there is a bit of a limitation when things start getting more complex.

	You've probably noticed that, unlike the method before, we've had to hardcode the time format in the timeHandler function. What happens when you want to pass information or variables from main() to a handler?

	A neat approach is to put our handler logic into a closure, and close over the variables we want to use, like this:

	An Introduction to Handlers and Servemuxes in Go – Alex Edwards
	https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func timeHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm + "\n"))
	}
}

func main() {
	mux := http.NewServeMux()

	// th := timeHandler(time.RFC1123)
	// mux.Handle("/time", th)
	mux.HandleFunc("/time", timeHandler(time.RFC1123))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}
