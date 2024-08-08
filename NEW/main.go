package main

import (
	"fmt"
)

func main() {
	str := "   hey there   we are   good     "

	word := SplitS(str)
	result := ""
	for i, ch := range word {
		if i != len(word) -1 {
			result += string(ch) + " "
		} else {
			result += string(ch)
		}
	}
	fmt.Println(result)

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
	// fmt.Println(word)
	return word
}
