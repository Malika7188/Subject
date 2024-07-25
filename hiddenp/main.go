package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	if s1 == "" {
		z01.PrintRune('1')
		z01.PrintRune('\n')
	}
	j := 0

	for i := 0; i < len(s2); i++ {
		if s2[i] == s1[j] {
			j++

			if j == len(s1) {
				z01.PrintRune('1')
				z01.PrintRune('\n')
				return
			}
		}
	}
	z01.PrintRune('0')
	z01.PrintRune('\n')
}
