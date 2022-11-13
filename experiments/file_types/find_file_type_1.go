/*
	go run find_file_type_1.go -i <file-path>

	*.json -> text/plain; charset=utf-8
	*.yml -> text/plain; charset=utf-8
	*.png -> image/png
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("starting main() function")

	inputPath := flag.String("i", "", "accepts relative or absolute paths to a yaml, json, or database input")
	flag.Parse()
	fmt.Printf("detected flag -i (%s)\n", *inputPath)

	checkFileType(*inputPath)
}

// rayrutjes/golang detect content type of a file -- https://gist.github.com/rayrutjes/db9b9ea8e02255d62ce2
func checkFileType(inputFilePath string) (string, error) {
	// open a file
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("failed to open file %q\n", inputFilePath)
		panic(err)
	}

	// defer closing file till function returns output
	defer file.Close()

	// Golang ReadSeeker Examples, io.ReadSeeker Golang Examples - HotExamples -- https://golang.hotexamples.com/examples/io/-/ReadSeeker/golang-readseeker-function-examples.html
	// How do I go from io.ReadCloser to io.ReadSeeker? - Stack Overflow -- https://stackoverflow.com/questions/37718191/how-do-i-go-from-io-readcloser-to-io-readseeker#:~:text=io.ReadSeeker%20is%20the%20interface,int)%20(int64%2C%20error)
	seeker := io.ReadSeeker(file)

	// Only the first 512 bytes are used to sniff the content type -- https://golang.org/src/net/http/sniff.go?s=646:688#L11
	buffer := make([]byte, 512)

	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	_, err = seeker.Seek(0, io.SeekStart)
	if err != nil {
		return "no bytes found", err
	}

	bytesRead, err := seeker.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	// Slice to remove fill-up zero values which cause a wrong content type detection in the next step
	buffer = buffer[:bytesRead]
	fmt.Println(buffer)

	contentType := http.DetectContentType(buffer)
	fmt.Println(contentType)

	return contentType, nil
}
