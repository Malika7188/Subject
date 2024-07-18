package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	str3 := os.Args[3]

	 result := ""
	for _, char := range str1 {
		for _, char1 := range str2 {
			if char == char1 {
				result += str3
			} else {
				result += string(char)
			}
		}
	}
	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}