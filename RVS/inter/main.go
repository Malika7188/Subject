package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	seen := make(map[rune]bool)

	for _, char := range str1 {
		for _, char1 := range str2 {
			if char == char1 {
				seen[char] = true
			}
		}
	}
	printed := make(map[rune]bool)

	for _, char := range str1 {
		if seen[char] && !printed[char] {
			z01.PrintRune(char)
			printed[char] = true
		}
	}
	z01.PrintRune('\n')
}
