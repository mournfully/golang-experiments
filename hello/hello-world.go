// https://gobyexample.com/hello-world
// use `go run` to run a program
// use `go build` to build the binary

package main

import {
	"fmt"
	"math"
}

func main() {
	// substrings can be added together with '+'
	fmt.Println("hello" + "go" + "lang")

	// fairly standard integers, floats, and booleans.
	fmt.Println("6+9 =", 6+9)
	fmt.Println("8.0/4.0 =", 8.0/4.0)
	fmt.Println(true || false)

	// variables, I guess no more `let`
	var a = "initial"
	f := "apple" // `:=` is shorthand for var
	var b, c int = 1, 2
	var e int // defaults to 0
}
