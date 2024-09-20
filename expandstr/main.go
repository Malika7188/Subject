package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	str := os.Args[1]
	words := SplitS(str)
	res := ""

	for i := 0; i < len(words); i++ {
		if len(words)-1 != ' ' {
			res += words[i] + "   "
		} else {
			res += words[i]
		}
	}
	// for _, char := range words {
	// 	if char == " " && len(words)-1 == ' ' {
	// 		res += string(char)
	// 	} else {
	// 		res += string(char) + " "
	// 	}
	// }
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

// res := ""
// str := os.Args[1]
// if len(str) == 0 {
// 	fmt.Println()
// 	return
// }
// arr := []string{}
// for i, char := range str {
// 	if char != ' ' {
// 		result += string(char)
// 	}
// 	if (char == ' ' && result != "") || (i == len(str)-1 && result != "") {
// 		arr = append(arr, result)
// 		result = ""
// 	}

// }
// for i, v := range arr {
// 	if i != len(arr)-1 {
// 		fmt.Print(v + " ")
// 	} else {
// 		fmt.Print(v)
// 	}
// }
// fmt.Println()
