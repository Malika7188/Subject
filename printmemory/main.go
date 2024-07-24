package main

import (
	"unicode"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	res := []string{}

	for _, char := range arr {
		if !unicode.IsGraphic(rune(char)) {
			res = append(res, "00")
		} else {
			res = append(res, Hex(int(char)))
		}
	}
	for i := 0; i < len(res); i++ {
		PrintS(res[i])
		z01.PrintRune(' ')
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for _, ch := range arr {
		if ch < 32 || ch > 128 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(ch))
		}
	}
	z01.PrintRune('\n')

}
func PrintS(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
func Hex(num int) string {
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdefgh"

	result := ""
	for num > 0 {
		digit := num % 16
		result = string(hex[digit]) + result
		num /= 16
	}
	return result
}
