package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	for i, c := range arr {
		if !unicode.IsGraphic(rune(c)) {
			z01.PrintRune('0')
			z01.PrintRune('0')
			if i == 3 || i == 7 || i == 9 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(' ')
			}
		} else {
			num := int(c)
			hexChars := "0123456789ABCDEF"
			result := ""
			for num > 0 {
				remainder := num % 16
				result = string(hexChars[remainder]) + result
				num = num / 16
			}
			if i == 3 || i == 7 || i == 9 {
				PrintStr(result)
				z01.PrintRune('\n')
			} else {
				PrintStr(result)
				z01.PrintRune(' ')
			}
		}
	}

	for _, c := range arr {
		if unicode.IsGraphic(rune(c)) {
			z01.PrintRune(rune(c))
		} else {
			z01.PrintRune('.')
		}
	}

	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
