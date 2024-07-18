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

	val := Atoi(args)

	hex := "0123456789abcdef"
	result := ""
	for val > 0 {
		digit := val%16
		result = string(hex[digit]) + result
		val /= 16
	}
	fmt.Println(result)


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
	result := 0
	for _, val := range str {
		if val < '0' || val > '9' {
			os.Exit(0)
		}
		result = (result * 10) + int(val-'0')
	}
	if isneg == true {
		result *= -1
	}
	return result
}
