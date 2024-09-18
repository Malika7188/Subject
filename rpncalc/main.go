package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	args := Split(os.Args[1])

	num, err := rpn(args)
	if err != "" {
		fmt.Println("Error")
		return
	}
	fmt.Println(num)
}
func rpn(s []string) (int, string) {
	number := []int{}

	for _, v := range s {
		if v == "+" || v == "-" || v == "*" || v == "%" || v == "/" {
			if len(number) < 2 {
				return 0, "Error"
			}
			b := number[len(number)-1]
			a := number[len(number)-2]
			number = number[:len(number)-2]
			cal := 0
			if v == "+" {
				cal = a + b
			}
			if v == "-" {
				cal = a - b
			}
			if v == "*" {
				cal = a * b
			}
			if v == "%" {
				cal = a % b
			}
			if v == "/" {
				cal = a / b
			}
			number = append(number, cal)
		} else {
			num, err := strconv.Atoi(v)
			if err != nil {
				return 0, "Error"
			}
			number = append(number, num)
		}
	}
	if len(number) != 1 {
		return 0, "Error"
	}
	return number[0], ""
}
func Split(s string) []string {
	word := []string{}
	res := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			res += string(s[i])
		} else if s[i] == ' ' && res != "" {
			word = append(word, res)
			res = ""
		}
	}
	if res != "" {
		word = append(word, res)
	}
	return word
}
