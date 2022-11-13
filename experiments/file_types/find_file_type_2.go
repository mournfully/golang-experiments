/*
	trying to make a flag input agnostic might not be such a bright idea - i'll just have different flags for different input types

	go run find_file_type_2.go <file-path>

	*.json -> text/plain; charset=utf-8
	*.yml -> text/plain; charset=utf-8
	*.png -> image/png

	---

	The standard library's code is only supposed to detect certain types (like HTML that contains one of a few common tags) according to a certain standardized algorithm--more in https://golang.org/src/net/http/sniff.go. You could use github.com/rakyll/magicmime, which uses libmagic and cgo to guess at more file types.

	go - Golang “net/http” DetectContentType error - Stack Overflow
	https://stackoverflow.com/questions/40601725/golang-net-http-detectcontenttype-error
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// golang detect content type of a file -- https://gist.github.com/hkak03key/06b25a3f4f0bbd8d23d361fa8eb0dff8
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer[:n])
	fmt.Println(contentType)
}
