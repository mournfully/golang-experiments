/*
	This time around, instead of banging my head against my keyboard when I get stuck. I decided to look at the solution briefley and then try and understand how they'd come to that conlclusion from the official documentation. I'm not sure if this is a better method...

	-[x] url redirection
	Test with `go run main.go` and then on your browser go to `http://localhost:8080/` - was confused for a bit lol

	-[x] parse map pathsToUrls and redirect from entries there

	-[ ] create a cli flag to set path to input source
		- yaml
		- json
		- database
		- and warn user if no accepted methods found

	-[ ] parse from yaml file
	-[ ] convert data to a map
	-[ ] reuse maphandler() to parse map

	-[ ] dockerize environment
*/

package main

import (
	"exercise-2-urlshort/urlshort"
	"fmt"
	"net/http"
)

func main() {
	// I think it makes sense to send this through to mapHandler() as fallback, because we're only writing to here if we can't redirect which doesn't need either reads or writes :O
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	// Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", failHandler)
	return mux
}

func failHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "couldn't redirect - given path doesn't match any stored urls")
}
