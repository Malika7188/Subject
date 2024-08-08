package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		return
	}
	args := os.Args[1:]

	word := ""
	for _, str := range args {
		for _, ch := range str {
			if ch >= 'A' && ch <= 'Z' {
				ch += 32
			}
			word += string(ch)
		}
		word += " "
	}
	res := ""
	for i := 0; i < len(word)-1; i++ {
		ch := word[i]
		// if IsDig(rune(word[i])) {
		// 	return
		// }
		if ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' {
			if i == len(word)-1 || i < len(word)-1 && (word[i+1]) == ' ' {
				ch -= 'a'-'A'
			}
			res += string(ch)
		} else if ch == ' ' {
			res += string(ch)
		}
	}
	fmt.Print(res)
	fmt.Println()
}
func IsDig(s rune) bool {
	return s >= '0' && s <= '9'
}
