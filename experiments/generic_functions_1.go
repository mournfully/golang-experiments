/*
	This is how you would apporach such a problem before generics

	go - How do you make a function accept multiple types? - Stack Overflow
	https://stackoverflow.com/questions/40145569/how-do-you-make-a-function-accept-multiple-types

	Tutorial: Getting started with generics - The Go Programming Language
	https://go.dev/doc/tutorial/generics#add_generic_function
*/

package main

import "fmt"

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	// Call the two functions declared below to find the sum of each map’s values and print the result.
	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

}

// SumInts takes a map of string to int64 values and adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats takes a map of string to float64 values and adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
