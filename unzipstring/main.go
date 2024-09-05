package main

import (
	"fmt"
	"strconv"
)

func Unzipstring(s string) string {
	res := ""
	if !Valid(s) {
		return "invalid Input"
	}
	if s == "" {
		return "ivalid input"
	}
	result := []rune{}

	for _, c := range s {
		if isDig(c) {
			res += string(c)
		} else if isLetter(c) {
			count, _ := strconv.Atoi(res)
			res = ""

			for i := 0; i < count; i++ {
				result = append(result, c)
			}
		}
	}
	r := ""
	for _, c := range result {
		r += string(c)
	}
	return r
}

func Valid(s string) bool {
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if i%2 == 0 {
			if !isDig(c) {
				return false
			}
		} else {
			if !isLetter(c) {
				return false
			}
		}

	}
	return true
}
func isDig(s rune) bool {
	return s >= '0' && s <= '9'
}
func isLetter(s rune) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}
func main() {
	fmt.Println(Unzipstring("2f"))
	fmt.Println(Unzipstring("2O5u2H0e"))
	fmt.Println(Unzipstring("3p6i1W"))
	fmt.Println(Unzipstring("6H8a"))
	fmt.Println(Unzipstring("2p4;7g"))
	fmt.Println(Unzipstring("2a 6p8f"))
	fmt.Println(Unzipstring("2t4dD"))
	fmt.Println(Unzipstring("82t4D"))
	fmt.Println(Unzipstring(""))
}
