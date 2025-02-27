package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	res := ""
	seen := make(map[rune]bool)
	// seen := false
	for _, c := range s1 + s2 {
		if !seen[c] {
			seen[c] = true
			res += string(c)
		}
	}
	fmt.Println(res)
	// for _, c := range s1 {
	// 	for _, ch := range s2 {
	// 		if c != ch {
	// 			seen[c] = true
	// 		}
	// 	}
	// }
	// print := make(map[rune]bool)
	// for _, c := range s1 {
	// 	if seen[c] && !print[c] {
	// 		z01.PrintRune(c)
	// 		print[c] = true
	// 	}
	// }
	// z01.PrintRune('\n')
}
