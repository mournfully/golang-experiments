package urlshort

import (
	"fmt"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v3"
)

/*  FallbackHandler() will be what the webserver will default to if it can't find any valid paths
 */
func FallbackHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("couldn't redirect - path '%s' doesn't match any stored urls", r.URL.Path)
	fmt.Fprintln(w, message)
}

/* MapHandler() will return an http.HandlerFunc that will attempt to
 * 	- map any paths (localhost:8080/<path>) to corresponding URL (https://<url>)
 *	- if provided path =/= in map, then do fallback http.Handler
 */
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	// Introduction to HTTP with Go : Our first microservice - YouTube -- https://www.youtube.com/watch?v=MKkokYpGyTU
	// docker-development-youtube-series/golang/introduction/part-3.http at master · marcel-dempers/docker-development-youtube-series -- https://github.com/marcel-dempers/docker-development-youtube-series/tree/master/golang/introduction/part-3.http
	// http package - net/http - Go Packages -- https://pkg.go.dev/net/http
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Golang Maps/HashMap/HashTable - ULTIMATE Golang Tutorial! - YouTube -- https://www.youtube.com/watch?v=92Q8n3LlMOY
		url, ok := pathsToUrls[path]
		if ok {
			/* found, two methods for redirection in official docs.
			 *  - http.RedirectHandler(), no error but had no noticeable effect
			 *  - http.Redirect(), worked for some reason
			 */
			// HTTP response status codes - HTTP | MDN -- https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#redirection_messages
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			// fmt.Println("couldn't redirect - given path doesn't match any stored urls")
			fallback.ServeHTTP(w, r)
		}
	}
}

/*  YamlHandler() will take a yaml input's file path and send data to
 *  various subfunctions in order to eventually send it all to MapHandler()
 */
func YamlHandler(inputFilePath string, fallback http.Handler) (http.HandlerFunc, error) {
	parsedData, err := readFile(inputFilePath)
	if err != nil {
		return nil, err
	}
	mappedData := mapBuilder(parsedData)
	return MapHandler(mappedData, fallback), nil
}

/*  mapBuilder() will take in any parsedData (in struct format)
 *  and transform it into something a hashmap can accept
 */
func mapBuilder(parsedData []yamlDataFormat) map[string]string {
	mappedData := make(map[string]string)
	for _, entry := range parsedData {
		mappedData[entry.Path] = entry.URL
	}
	return mappedData
}

/*  readFile() takes in a file path and outputs data
 *  in a predefined golang struct to send to mapBuilder()
 */
func readFile(inputFilePath string) ([]yamlDataFormat, error) {
	// attempt to open file
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("failed to open file %q\n", inputFilePath)
		return nil, err
	}

	// defer closing file till function is done
	defer file.Close()

	// parse input file into a slice of bytes
	// Go: Read a whole file into a string (byte slice) | Programming.Guide -- https://programming.guide/go/read-file-to-string.html
	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Printf("failed to parse provided file %q\n", inputFilePath)
		return nil, err
	}

	// force unorganized input data into a predefined struct
	var parsedData []yamlDataFormat
	err = yaml.Unmarshal(content, &parsedData)
	if err != nil {
		return nil, err
	}
	return parsedData, nil
}

type yamlDataFormat struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
