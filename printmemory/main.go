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
	for i, cha := range arr {
		Hex(int(cha))
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != len(arr)-1 {
			z01.PrintRune(' ')
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

// func main() {
// 	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
// }

func PrintStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
func Hex(num int) {
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
	if len(res) == 1 {
		res = "0" + res
	}
	PrintStr(res)
}
