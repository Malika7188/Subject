package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	if args == "" {
		fmt.Println()
		return
	}
	str := Split(args)
	word := ""

	for i := len(str) - 1; i >= 0; i-- {
		if str[len(str)-1] != " " && i != 0 {
			word += str[i] + " "
		} else {
			word += str[i]
		}
	}
	fmt.Println(word)
}
func Split(str string) []string {
	res := ""
	word := []string{}

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			res += string(str[i])
		} else if str[i] == ' ' && res != "" {
			word = append(word, res)
			res = ""
		}
	}
	if res != "" {
		word = append(word, res)
	}
	return word
}
