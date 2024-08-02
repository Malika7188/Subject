package main

import "fmt"

func WordFlip(str string) string {

	if str == "" {
		return "Invalid Output\n"
	}
	res := Split(str)

	if len(res) == 0 {
		return "\n"
	}
	resul := ""

	for i := len(res) - 1; i >= 0; i-- {
		// resul += res[i]
		// if i != 0 {
		// 	resul += " "
		// }
		if res[len(res[i])-1] == " " {
			resul += string(res[i])
		} else {
			resul += string(res[i])  + " "
		}
	}
	return resul + "\n"
}

func Split(s string) []string {
	word := []string{}
	result := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			result += string(s[i])
		} else if s[i] == ' ' && result != "" {
			word = append(word, result)
			result = ""
		}
	}
	if result != "" {
		word = append(word, result)
	}
	return word
}
func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
