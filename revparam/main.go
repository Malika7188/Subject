package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 1 {
		return
	}
	args := os.Args[1:]
	result := ""
	for i := len(args) - 1; i >= 0; i-- {
		result += args[i] + "\n"
	}
	PrintS(result)
	z01.PrintRune('\n')
}

func PrintS(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
