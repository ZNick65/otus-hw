package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(reverseString("Hello, OTUS!"))
}

// Reverse returns its argument string reversed rune-wise left to right.
func reverseString(greeting string) string {
	return stringutil.Reverse(greeting)
}
