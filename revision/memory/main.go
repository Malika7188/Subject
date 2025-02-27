package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	res := []string{}

	for _, c := range arr {
		if c == 0 {
			res = append(res, "00")
		} else {
			res = append(res, Hex(int(c)))
		}
	}
	// fmt.Println(res)
	for i := 0; i < len(res); i++ {
		P(res[i])

		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for _, c := range arr {
		if c < 32 || c > 126 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(c))
		}
	}
	z01.PrintRune('\n')
}

func Hex(n int) string {
	res := ""
	hexV := "0123456789abcdef"

	for n > 0 {
		dig := n % 16
		res = string(hexV[dig]) + res
		n /= 16
	}
	if len(res) != 2 {
		res = "0" + res
	}
	return res
}
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
func P(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
