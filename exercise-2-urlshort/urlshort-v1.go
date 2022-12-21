//go:build exclude

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
