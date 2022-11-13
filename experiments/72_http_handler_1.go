/*
Processing HTTP requests with Go is primarily about two things: handlers and servemuxes.

Whereas a servemux (also known as a router) stores a mapping between the predefined. URL paths for your application and the corresponding handlers.
Usually you have one servemux for your application containing all your routes.

Go's net/http package ships with the simple but effective http.ServeMux servemux, plus a few functions to generate common handlers including http.FileServer(), http.NotFoundHandler() and http.RedirectHandler().

An Introduction to Handlers and Servemuxes in Go – Alex Edwards
https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
*/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to create an empty servemux.
	mux := http.NewServeMux()

	// Use the http.RedirectHandler() function to create a handler which 307 redirects all requests it receives to https://google.com.
	rh := http.RedirectHandler("https://google.com", 307)

	// Next we use the mux.Handle() function to register this with our new servemux, so it acts as the handler for all incoming requests with the URL path /google.
	mux.Handle("/google", rh)

	// Then we create a new server and start listening for incoming requests with the http.ListenAndServe() function, passing in our servemux for it to match requests against as the second parameter
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
	/*
		$ curl -IL localhost:8080/google
		HTTP/1.1 307 Temporary Redirect

		$ curl -IL localhost:8080/foo
		HTTP/1.1 404 Not Found
	*/

}
