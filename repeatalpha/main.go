package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// if len(os.Args) != 1 {
	// 	return
	// }
	args := os.Args[1]

	count := 0

	for _, char := range args {
		if char >= 'a' && char <= 'z' {
			count = int(char - 'a' + 1)
		} else if char >= 'A' && char <= 'Z' {
			count = int(char - 'A' + 1)
		}
		for i := 0; i < count; i++ {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
