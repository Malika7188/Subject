package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]
	result := []rune{}
	for _, char := range word {
		if char >= 'a' && char <= 'z' {
			char = (char - 'a'+13)%26 + 'a'
			result = append(result, char)
		} else if char >= 'A' && char <= 'Z' {
			char = (char - 'a'+13)%26 + 'a'
			result = append(result, char)
		}
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
