package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}
	for _, char := range args[0] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
