package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1:]

	for _, c := range arg {
		res := Cap(c)
		fmt.Println(res)
	}


}
func Cap(s string) string {
	res := ""

	for i, c := range s {
		if i != len(s)-1 {
			if Up(c) && s[i+1] != ' ' {
				res += string(c + 32)
			} else if low(c) && s[i+1] == ' ' {
				res += string(c - 32)
			} else {
				res += string(c)
			}
		} else {
			if i == len(s)-1 && low(c) {
				res += string(c - 32)
			} else {
				res += string(c)
			}
		}
	}
	return res
}
func Up(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
func low(r rune) bool {
	return r >= 'a' && r <= 'z'
}
