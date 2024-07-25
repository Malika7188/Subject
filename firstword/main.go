package main

import "fmt"

func FirstWord(s string) string {
	res := ""
	word := false

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			res += string(s[i])
			word = true
		} else if word {
			break
		}
	}
	return res + "\n"
}
func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
