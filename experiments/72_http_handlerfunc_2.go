/*
	The handlers that ship with net/http are useful, but most of the time when building a web application you'll want to use your own custom handlers instead. The first thing to explain is that anything in Go can be a handler so long as it satisfies the http.Handler interface, which looks like this:

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}

	An Introduction to Handlers and Servemuxes in Go – Alex Edwards
	https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

/*
	For simple cases (like the example below) defining new a custom type just to make a handler feels a bit verbose. Fortunately, we can rewrite the handler as a simple function instead, see page "_3"
*/

// A custom handler which responds with the current time in a specific format. Like this:
func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	mux := http.NewServeMux()

	// Initialise the timeHandler in exactly the same way we would any normal struct.
	th := timeHandler{format: time.RFC1123}

	// Like the previous example, we use the mux.Handle() fnction to register this with our ServeMux.
	mux.Handle("/time", th)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
	/*
		curl localhost:8080/time
		The time is: Sat, 12 Nov 2022 16:55:37 EST┌
	*/
}
