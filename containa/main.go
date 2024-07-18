package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	args := os.Args[1]
	// char := []rune{'a','A'}
	// seen := make(map[rune]bool)
	containA := false

	for _, char := range args {
		if char == 'a' {
			containA = true
			os.Stdout.WriteString("contains a")
		}
	}
	if !containA {
		os.Stdout.WriteString("Contains the letter")
	}
	z01.PrintRune('\n')
}
