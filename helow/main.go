package main

import "github.com/01-edu/z01"

func main() {
	args := "hello world"

	for _, char := range args{
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}