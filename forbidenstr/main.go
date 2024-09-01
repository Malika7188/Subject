package main

import (
	"fmt"
	"strings"
)

func forbid(s string, sub []string) bool {
	for i := 0; i < len(sub); i++ {
		if strings.Contains(s, sub[i]) {
			return true
		}
	}
	return false
}
func main() {
	Input, sub := "hello", []string{"ll", "xx"}
	fmt.Println(forbid(Input, sub))
}

// Check for Forbidden Substrings: Given a string and a slice of forbidden substrings, write a function that returns true if the string contains any of the forbidden substrings.

//     Example: Input: "hello", []string{"ll", "xx"}
//     Output: true
