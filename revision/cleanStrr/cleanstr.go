package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := Split(os.Args[1])
	res := ""

	for i, c := range str {
		if c != " " && i != len(str)-1 {
			res += c + " "
		} else {
			res += c
		}
	}
	fmt.Println(res)
}
func Split(s string) []string {
	res := ""
	word := []string{}

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			res += string(s[i])
		} else if s[i] == ' ' && res != "" {
			word = append(word, res)
			res = ""
		}
	}
	if res != "" {
		word = append(word, res)
	}
	return word
}
