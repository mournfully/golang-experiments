// https://gobyexample.com/hello-world
// use `go run` to run a program
// use `go build` to build the binary

package main

import (
	"fmt"
	"math"
)

func main() {
	// substrings can be added together with '+'
	fmt.Println("hello" + "go" + "lang")

	// fairly standard integers, floats, and booleans.
	fmt.Println("6+9 =", 6+9)
	fmt.Println("8.0/4.0 =", 8.0/4.0)
	fmt.Println(true || false)

	// variables, var <name> <type> = <value>
	// strings, integers, floats, booleans
	var a = "initial"
	f := "apple" // `:=` shorthand for `var string`
	var e int    // uninitialized = zero-valued
	var b, c int = 1, 2
	fmt.Printf("strings: %s %s | uninitialized: %d | integers: %d %d\n", a, f, e, b, c)
	// doesn't make a new line btw ^ so add \n

	// constants, how is this different from variables?
	const s string = "constant"
	const n = 500
	const d = 3e20 / n
	fmt.Println(d) // NUMERIC constants by default have no type
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
