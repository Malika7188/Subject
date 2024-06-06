package main

import "github.com/01-edu/z01"

func main() {
	s := "0 1 2 3 4 5 6 7 8 9"

	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}