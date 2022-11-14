package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

// if no valid paths are found, use this handler
func FallbackHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("couldn't redirect - path '%s' doesn't match any stored urls\n", r.URL.Path)
	fmt.Fprintln(w, message)
}

// if valid path is found return an http.HandlerFunc or use fallback
func MapHandler(pathsToUrls map[string]string, mux http.Handler) http.HandlerFunc {
	return func(write http.ResponseWriter, read *http.Request) {
		url, ok := pathsToUrls[read.URL.Path]
		if ok {
			http.Redirect(write, read, url, http.StatusPermanentRedirect)
		} else {
			mux.ServeHTTP(write, read)
		}
	}
}

type dataFormat struct {
	Path string
	URL  string
}

// take data in structure []dataFormat -> add each entry to hashmap pathsToUrls
func mapBuilder(parsedData []dataFormat) map[string]string {
	mappedData := make(map[string]string)
	for _, entry := range parsedData {
		mappedData[entry.Path] = entry.URL
	}
	return mappedData
}

// take yaml data in []bytes -> convert to []dataFormat ->  send to mapbuilder to convert to [string]string
func YamlHandler(yamlFileData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedData, err := yamlReader(yamlFileData)
	if err != nil {
		return nil, err
	}

	mappedData := mapBuilder(parsedData)
	return MapHandler(mappedData, fallback), nil
}

func yamlReader(yamlFileData []byte) (parsedData []dataFormat, err error) {
	err = yaml.Unmarshal(yamlFileData, &parsedData)
	return
}

// take json data in []bytes -> convert to []dataFormat ->  send to mapbuilder to convert to [string]string
func JsonHandler(jsonFileData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedData, err := jsonReader(jsonFileData)
	if err != nil {
		return nil, err
	}

	mappedData := mapBuilder(parsedData)
	return MapHandler(mappedData, fallback), nil
}

func jsonReader(jsonFileData []byte) (parsedData []dataFormat, err error) {
	err = json.Unmarshal(jsonFileData, &parsedData)
	return
}
