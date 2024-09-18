package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	words := SplitS(str)
	res := ""
	for _, char := range words {
		if char == " " {
			res += string(char)
		} else {
			res += string(char) + " "
		}
	}
	fmt.Println(res)
}

func SplitS(str string) []string {
	result := ""
	word := []string{}

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			result += string(str[i])
		} else if str[i] == ' ' && result != "" {
			word = append(word, result)
			result = ""
		}
	}
	if result != "" {
		word = append(word, result)
	}
	return word
}

