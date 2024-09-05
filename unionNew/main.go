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
	args := os.Args[1:]
	str1 := args[0]
	str2 := args[1]
	res := ""
	seen := make(map[rune]bool)
	for _, c := range str1 + str2 {
		if !seen[c] {
			seen[c] = true
			res += string(c)
		}
	}
	fmt.Println(res)
}

// 	words := ""

// 	for _, str := range os.Args[1:] {

// 		words += str
// 	}
// 	result := ""
// 	seen := make(map[rune]bool)
// 	for _, char := range words {
// 		if !seen[char] {
// 			seen[char] = true
// 			result += string(char)
// 		}
// 	}
// 	fmt.Println(result)
// }
