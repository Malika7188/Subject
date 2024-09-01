package main

import (
	"fmt"
	"strings"
)

func substr(s []string, sub string) []string {
	res := []string{}
	// seen := false
	for i := 0; i < len(s); i++ {
		if strings.Contains(string(s[i]), sub) {
			res = append(res, s[i])
		}
	}
	return res
}
func main() {
	Input, sub := []string{"hello", "world", "help", "hold"}, "he"
	fmt.Println(substr(Input, sub))
}

// Write a function that takes a slice of strings and a substring, returning a new slice of strings containing only those strings that have the substring.

//     Example: Input: []string{"hello", "world", "help", "hold"}, "he"
//     Output: []string{"hello", "help"}
