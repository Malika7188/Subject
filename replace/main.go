package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str1 := os.Args[1]
	str2 := os.Args[2]
	str3 := os.Args[3]

	replace(str1, str2, str3)
}

func replace(str1, str2, str3 string) {
	newchar := ""

	for _, char := range str1 {
		for _, char1 := range str2 {
			if char == char1 {
				newchar += str3
			} else {
				newchar += string(char)
			}
		}
	}
	for _, char := range newchar {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
