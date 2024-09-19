package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]

	slice := args[0]
	slice2 := Split(args[1], " ")

	if len(slice) == 0 {
		return
	}
	if len(slice) < 2 || slice[0] != '(' || slice[len(slice)-1] != ')' {
		return
	} else {
		slice = slice[1 : len(slice)-1]
	}

	slice1 := Split(slice, "|")
	count := 0
	for _, str := range slice2 {
		for _, s := range slice1 {
			if StringContains(str, s) {
				count++
				fmt.Printf("%v: %v\n", count, str)
				break
			}
		}
	}
}

func Split(s string, sep string) []string {
	var result []string
	var word string

	sepLen := len(sep)
	if sepLen == 0 {
		return []string{s}
	}

	for i := 0; i < len(s); i++ {
		if i+sepLen <= len(s) && s[i:i+sepLen] == sep {
			if word != "" {
				result = append(result, word)
				word = ""
			}
			i += sepLen - 1 // Skip over the separator
		} else {
			word += string(s[i])
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}

// Improved StringContains function to handle substring search
func StringContains(s string, r string) bool {
	if len(r) == 0 || len(s) == 0 {
		return false
	}

	// Search for the substring `r` in `s`
	for i := 0; i <= len(s)-len(r); i++ {
		if s[i:i+len(r)] == r {
			return true
		}
	}

	return false
}
