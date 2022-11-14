/*
	Solution:
	https://github.com/gophercises/urlshort/blob/master/students/dennisvis/main.go
	https://github.com/gophercises/urlshort/blob/master/students/dennisvis/urlshort/handler.go

	This time around, instead of banging my head against my keyboard when I get stuck. I decided to look at the solution briefley and then try and understand how they'd come to that conlclusion from the official documentation. I'm not sure if this is a better method...

	- [x] url redirection, test with `go run main.go` and then on your browser go to `http://localhost:8080/` - was confused for a bit lol

	- [x] parse map pathsToUrls and redirect from entries there

	- [x] create a cli flag to use yaml input file
	- [x] parse from yaml file
	- [x] convert data to a map
	- [x] reuse maphandler() to parse map

	- [x] parse from json file too

	- [ ] dockerize environment?
	- [ ] high availability redis database?
	- [ ] read from database instead of map
	- [ ] improve your program so that it's comparable to [ptman/urlredir: Educational URL redirector service in Go](https://github.com/ptman/urlredir)?
*/

package main

import (
	"bytes"
	"exercise-2-urlshort/handlers"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var (
	inputFilePath = flag.String("i", "", "accepts relative or absolute paths to a yaml/json/database file")
	exposedPort   = ":8080"
)

func getFileBytes(inputFilePath string) []byte {
	// open a file
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("failed to open file %q\n", inputFilePath)
		panic(err)
	}

	// parse a file
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		fmt.Printf("failed to parse file %q\n", inputFilePath)
	}

	return buffer.Bytes()
}

func main() {
	// initialize variables
	flag.Parse()
	mux := defaultMux()

	var handler http.Handler
	var err error
	extension := filepath.Ext(*inputFilePath)

	// input checking and calling appropriate functions
	switch extension {
	case ".yml":
		handler, err = handlers.YamlHandler(getFileBytes(*inputFilePath), mux)
		if err != nil {
			panic(err)
		}
	case ".json":
		handler, err = handlers.JsonHandler(getFileBytes(*inputFilePath), mux)
		if err != nil {
			panic(err)
		}
	case ".db":
		fmt.Printf("database functionality isn't up yet\n")
		os.Exit(1)
	default:
		fmt.Printf("input file's extension needs to be either yaml (.yml), json (.json), or a redis database (.db).\n")
		os.Exit(1)
	}

	fmt.Printf("Starting the server on %s\n", exposedPort)
	http.ListenAndServe(exposedPort, handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.FallbackHandler)
	return mux
}
