package main

import (
	// student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
	"github.com/01-edu/z01"
)

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

func PrintMemory(arr [10]byte) {
	result := []string{}
	for _, ch := range arr {
		if ch == 0 {
			result = append(result, "00")
		} else {
			s := Hex(int(ch))
			if len(s) == 1 {
				s = "0" + s
			}
			result = append(result, s)
		}

	}
	for i := 0; i < len(result); i++ {
		Prints(result[i])
		// z01.PrintRune('\n')
		// if i == 3 || i == 7 || i == 9 {
		// 	z01.PrintRune('\n')
		// }
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != len(result)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
	for _, ch := range arr {
		if ch >= 32 && ch <= 126 {
			z01.PrintRune(rune(ch))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')

}
func Hex(num int) string {
	res := ""
	vals := "0123456789abcdef"

	if num == 0 {
		return "0"
	}

	for num > 0 {
		dig := num % len(vals)
		res = string(vals[dig]) + res
		num /= len(vals)
	}
	return res
}

// func main() {
// 	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
// }
func Prints(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}
