// ~/golang/experiments/multiple_go_files/file-one.go
package main

import (
	"fmt"
)

func main() {
	funcFromFileTwo()
	fmt.Println("Hi from file one")
}
