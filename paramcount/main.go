package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// args := os.Args[1:]
	// count := len(args)

	// var result []rune

	// if count == 0 {
	// 	z01.PrintRune('0')
	// }
	// for count > 0 {
	// 	digit := count % 10
	// 	result = append([]rune{rune(digit + '0')}, result...)
	// 	count /= 10
	// }
	// for _, char := range result {
	// 	z01.PrintRune(char)
	// }
	// z01.PrintRune('\n')
	args := os.Args[1:]

	count := len(args)

	if count == 0 {
		z01.PrintRune('0')
	}

	result := []rune{}

	for count > 0 {
		digit := count%10
		result = append([]rune{rune(digit + '0')}, result...)
		count /= 10
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
