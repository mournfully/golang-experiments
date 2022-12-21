package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"example.com/handlers"
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
