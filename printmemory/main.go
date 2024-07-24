package main

import (
	// "unicode"

	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	for i, cha := range arr {
		Hex(int(cha))
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')

	for _, char := range arr {
		if char >= 32 && char <= 128 {
			PrintStr(string(char))
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
func Hex(num int){
	res := ""
	vals := "0123456789abcdef"

	if num == 0 {
		res = "00" + res
	}
	for num > 0 {
		digit := num % 16
		res = string(vals[digit]) + res
		num /= 16
	}
	PrintStr(res + " ")
}
