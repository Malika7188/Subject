package main

import "fmt"

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	res := SplitStr(s)
	for _, ch := range res {
		for i, w := range ch {
			if i == 0 && (w >= 'a' && w <= 'z') {
				return false
			}
		}
	}
	return true

}
func SplitStr(str string) []string {
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
	// fmt.Println(word)
	return word
}
func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

