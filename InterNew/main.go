package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]

	str1 := args[0]
	str2 := args[1]

	seen := make(map[rune]bool)

	for _, c := range str1 {
		for _, ch := range str2 {
			if c == ch {
				seen[c] = true
				// break
			}
		}
	}
	printed := make(map[rune]bool)
	for _, c := range str1 {
		if seen[c] && !printed[c] {
			printed[c] = true
			fmt.Print(string(c))
		}
	}
	fmt.Println()
	// res := ""
	// for _, c := range str1 {
	// 	ch := string(c)
	// 	if strings.Contains(str2, ch) && !seen[c] {
	// 		seen[c] = true
	// 		res += string(c)
	// 	}
	// }
	// fmt.Println(res)
}
