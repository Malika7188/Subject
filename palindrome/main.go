package main

import "fmt"

func Palindrome(s string) bool {
	res := ""

	for _, c := range s {
		if c != ' ' {
			if c >= 'A' && c <= 'Z' {
				c += 32
			}
			res += string(c)			
		}
	}

	result := ""
	for i := len(res) - 1; i >= 0; i-- {
		result += string(res[i])
	}
	if res == result {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(Palindrome("Racecar"))
	fmt.Println(Palindrome("hello"))
	fmt.Println(Palindrome("123321"))
}
