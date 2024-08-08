package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	str3 := os.Args[3]

	replace(str1, str2, str3)
	// seen := make(map[rune]bool

	// for _, ch := range str {
	// 	words += string(ch)
	// }
	// res := ""
	// for _, ch := range words {
	// 	if !seen[ch] {
	// 		seen[ch] = true
	// 		res += string(ch)
	// 	}
	// }
	// fmt.Println(res)

	// for _, ch := range str1 {
	// 	for _, char := range str2 {
	// 		if ch != char {
	// 			seen[ch] = true
	// 		}
	// 	}
	// }
	// printed := make(map[rune]bool)
	// for _, ch := range str1 {
	// 	if seen[ch] && !printed[ch] {
	// 		z01.PrintRune(ch)
	// 		printed[ch] = true
	// 	}
	// }
	// z01.PrintRune('\n')
}
func replace(str1, str2, str3 string) {
	res := ""

	for _, ch := range str1 {
		for _, char := range str2 {
			if ch == char {
				res += str3
			} else {
				res += string(ch)
			}
		}
	}
	fmt.Println(res)
}
