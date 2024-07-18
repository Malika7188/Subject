package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]

	result := ""

	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			char = ('z' - char + 'a')
			result += string(char)
		} else if char >= 'A' && char <= 'A' {
			char = ('Z' - char + 'A')
			result += string(char)
		} else {
			result += string(char)
		}
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
