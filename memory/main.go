package main

import "github.com/01-edu/z01"

func PrintM(arr [10]byte) {
	res := []string{}

	for _, char := range arr {
		if char == 0 {
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
		z01.PrintRune('\n')
	}
	for _, char := range arr {
		if char < 32 || char > 128 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(char))
		}
	}
}
func Hex(num int) string {
	if num == 0 {
		return ""
	}
	hexVal := "123456789abcdef"

	res := ""

	for num > 0 {
		digit := num % 16
		res = string(hexVal[digit]) + res
		num /= 16
	}
	return res
}
func PrintS(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
func main() {
	PrintM([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
