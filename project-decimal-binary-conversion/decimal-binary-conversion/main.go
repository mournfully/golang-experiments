package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Wrapper(934))
}

func IntToBinary(n int) (output []int) {
	// create array to store binary numbers
	var binary []int

	// counter for array
	for n > 0 {
		// store remainder in array
		binary = append(binary, n%2)
		n = n / 2
	}

	// print array in reverse order
	for j := len(binary) - 1; j >= 0; j-- {
		output = append(output, binary[j])
	}

	return output
}

func SliceToString(numbers []int) (out string) {
	string := make([]string, len(numbers))
	for i, v := range numbers {
		string[i] = strconv.Itoa(v)
	}
	return strings.Join(string, "")
}

func Wrapper(num int) (out string) {
	slice := IntToBinary(num)
	out = SliceToString(slice)
	return
}
