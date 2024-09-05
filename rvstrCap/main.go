package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1:]
	for _, str := range args {
		fmt.Println(st(str))
	}

}
func st(str string) string {
	res := ""
	for i, c := range str {
		if i != len(str)-1 {
			if Cap(c) && str[i+1] != ' ' {
				res += string(c + 32)
			} else if Low(c) && str[i+1] == ' ' {
				res += string(c - 32)
			} else {
				res += string(c)
			}
		} else {
			if i == len(str)-1 && Low(c) {
				res += string(c - 32)
			} else {
				res += string(c)
			}
		}

	}
	return res
}
func Cap(s rune) bool {
	return s >= 'A' && s <= 'Z'
}
func Low(s rune) bool {
	return s >= 'a' && s <= 'z'
}
