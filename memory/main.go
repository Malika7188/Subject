package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
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
		// z01.PrintRune('\n')
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}

	}
	// z01.PrintRune('\n')
	for _, char := range arr {
		if char < 32 || char > 126 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(char))
		}
	}
	z01.PrintRune('\n')
}

func Hex(num int) string {
	if num == 0 {
		return ""
	}
	hexVal := "0123456789abcdef"

	res := ""

	for num > 0 {
		digit := num % 16
		res = string(hexVal[digit]) + res
		num /= 16
	}
	if len(res) == 1 {
		res = "0" + res
	}
	return res
}
func PrintS(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}

// func main() {
// 	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
// }

func main() {
	table := [10]byte{}

	for j := 0; j < 5; j++ {
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = byte(random.IntBetween(13, 126))
		}
		challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table)
	}
	table2 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table2)
}
