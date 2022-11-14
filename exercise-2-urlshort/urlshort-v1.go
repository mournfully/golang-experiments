//go:build exclude

/*
	ptman/urlredir: Educational URL redirector service in Go
	https://github.com/ptman/urlredir

	This time around, instead of banging my head against my keyboard when I get stuck. I decided to look at the solution briefley and then try and understand how they'd come to that conlclusion from the official documentation. I'm not sure if this is a better method...

	- [x] url redirection
	Test with `go run main.go` and then on your browser go to `http://localhost:8080/` - was confused for a bit lol

	- [x] parse map pathsToUrls and redirect from entries there

	- [x] create a cli flag to use yaml input file
	- [x] parse from yaml file
	- [x] convert data to a map
	- [x] reuse maphandler() to parse map

	- [ ] parse from json file
	- [ ] read from database instead of map
	- [ ] dockerize environment
	- [ ] high availability redis database
*/

package main

import (
	"exercise-2-urlshort/urlshort"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	// initialize cli flags
	inputYamlPath := flag.String("i", "", "accepts relative or absolute paths to a yaml file")
	flag.Parse()

	// I think it makes sense to send this through to mapHandler() as fallback, because we're only writing to here if we can't redirect which doesn't need either reads or writes :O
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	// Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// What is the best way to test for an empty string in Go? - Stack Overflow -- https://stackoverflow.com/questions/18594330/what-is-the-best-way-to-test-for-an-empty-string-in-go
	if len(*inputYamlPath) > 0 {
		fmt.Printf("yaml input flag detected, reading... %s \n", *inputYamlPath)

		yamlHandler, err := urlshort.YamlHandler(*inputYamlPath, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)
	} else {
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", mapHandler)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", urlshort.FallbackHandler)
	return mux
}
