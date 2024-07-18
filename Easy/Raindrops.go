package main

import (
	"fmt"
	"strconv"
)

func Convert(number int) string {
	result := ""

	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(number) // Return the number as a string if no sounds
	}

	return result
}

func main() {
	// Test examples
	fmt.Println(Convert(28)) // Output: Plong
	fmt.Println(Convert(30)) // Output: PlingPlang
	fmt.Println(Convert(34)) // Output: 34
}
