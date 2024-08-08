package main

import (
	"fmt"
	"os"
	// "github.com/01-edu/z01"
	// "strconv"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	str := os.Args[1:]

	for _, ch := range str {
		words := Split(ch)
		for i, char := range words {
			words[i] = wordStr(char)
		}
		// res := join(words)
		fmt.Println(join(words))
	}

}

func wordStr(s string) string {
	if len(s) == 0 {
		return s
	}
	res := []rune(s)

	for i := 0; i < len(res); i++ {
		if i == len(res)-1 {
			res[i] = IsCap(res[i])
		} else {
			res[i] = IsLow(res[i])
		}
	}
	return string(res)
}
func IsCap(s rune) rune {
	if s >= 'a' && s <= 'z' {
		return s - 'a' + 'A'
	}
	return s
}
func IsLow(s rune) rune {
	if s >= 'A' && s <= 'Z' {
		return s - 'A' + 'a'
	}
	return s
}
func Split(s string) []string {
	var words []string
	start := -1
	for i, c := range s {
		if c != ' ' && start == -1 {
			start = i
		} else if c == ' ' && start != -1 {
			words = append(words, s[start:i])
			start = -1
		}
	}
	if start != -1 {
		words = append(words, s[start:])
	}
	return words
}
func join(s []string) string {
	res := ""
	for i, c := range s {
		if i > 0 {
			res += " "
		} else {
			res += c
		}
	}
	return res
}
// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		return
// 	}
// 	args := os.Args[1:]

// 	for _, arg := range args {
// 		words := splitIntoWords(arg)
// 		for i, word := range words {
// 			words[i] = processWord(word)
// 		}
// 		fmt.Println(joinWords(words))
// 	}
// }

// func splitIntoWords(s string) []string {
// 	var words []string
// 	start := -1
// 	for i, c := range s {
// 		if c != ' ' && start == -1 {
// 			start = i
// 		} else if c == ' ' && start != -1 {
// 			words = append(words, s[start:i])
// 			start = -1
// 		}
// 	}
// 	if start != -1 {
// 		words = append(words, s[start:])
// 	}
// 	return words
// }

// func processWord(word string) string {
// 	if len(word) == 0 {
// 		return word
// 	}

// 	runes := []rune(word)
// 	for i := 0; i < len(runes); i++ {
// 		if i == len(runes)-1 {
// 			runes[i] = toUpper(runes[i])
// 		} else {
// 			runes[i] = toLower(runes[i])
// 		}
// 	}
// 	return string(runes)
// }

// func toUpper(r rune) rune {
// 	if r >= 'a' && r <= 'z' {
// 		return r - 'a' + 'A'
// 	}
// 	return r
// }

// func toLower(r rune) rune {
// 	if r >= 'A' && r <= 'Z' {
// 		return r - 'A' + 'a'
// 	}
// 	return r
// }

// func joinWords(words []string) string {
// 	result := ""
// 	for i, word := range words {
// 		if i > 0 {
// 			result += " "
// 		}
// 		result += word
// 	}
// 	return result
// }
