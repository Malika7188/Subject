package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]

	num := Atoi(args)
	res := ""
	hexvals := "0123456789abcdef"

	for num > 0 {

		res = string(hexvals[num%16]) + res

		num /= 16
	}
	fmt.Println(res)
}

func Atoi(str string) int {
	isneg := false
	if str[0] == '-' {
		isneg = true
		str = str[1:]
	}
	if str[0] == '+' {
		isneg = false
		str = str[1:]
	}
	var result int

	for _, char := range str {
		if char >= '0' && char <= '9' {
			result = (result * 10) + int(char-'0')
		} else {
			return 0
		}
	}
	if isneg {
		result *= -1
	}
	return result
}
