package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
	// "github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
	}
	words := os.Args[1]
	res := split(words, ",")
	s := ""
	for i := 1; i < len(res); i++ {
		s += res[i] + " "

	}
	s += res[0]
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}

func split(str, sep string) []string {
	result := []string{}
	word := ""
	 
	for i := 0; i < len(str); i++ {
		if string(str[i]) != sep {
			word += string(str[i])
		} else if string(str[i]) == sep && word != "" {
			result = append(result, word)
			word = ""
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
