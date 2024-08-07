package main

import "github.com/01-edu/z01"

// func PrintCombN() {
// 	for a := '0'; a <= '7'; a++ {
// 		for b := a + 1; b <= '8'; b++ {
// 			for c := b + 1; c <= '9'; c++ {
// 				z01.PrintRune(a)
// 				z01.PrintRune(b)
// 				z01.PrintRune(c)
// 				if a == '7' && b == '8' && c == '9' {
// 					z01.PrintRune('\n')
// 				} else {
// 					z01.PrintRune(',')
// 					z01.PrintRune(' ')
// 				}
// 			}
// 		}
// 	}
// }

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i == '7' && j == '8' && k == '9' {
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}

}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}
