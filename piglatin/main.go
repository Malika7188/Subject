package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		return
	}
	str := args[1]

	result := PigLatiin(str)

	if result == "" {
		fmt.Println("No Vowels")
	} else {
		fmt.Println(result)
	}
}

func PigLatiin(str string) string {

	for i, ch := range str {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			if i == 0 {
				return str + "ay"
			}
			return str[i:] + str[:i] + "ay"
		}
	}
	return ""
}
