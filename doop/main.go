package main

import (
	//"errors"
	//"log"
	"os"
	//"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	val1 := Atoi(os.Args[1])
	val2 := Atoi(os.Args[3])

	operator := os.Args[2]
	operators := []string{"+", "-", "/", "*", "%"}

	op := false

	for _, value := range operators {
		if operator == value {
			op = true
			break
		}
	}
	if !op {
		return
	}
	var answer int

	if operator == "%" && val2 == 0 {
		os.Stdout.WriteString("No modulo by 0\n")
		return
	}
	if operator == "/" && val2 == 0 {
		os.Stdout.WriteString("No division by 0\n")
		return
	}
	if val1 >= 9223372036854775807 || val1 <= -9223372036854775807 {
		return
	}
	if val2 <= -9223372036854775807 || val2 >= 9223372036854775807 {
		return
	}

	if operator == "+" {
		answer = val1 + val2
	}
	if operator == "-" {
		answer = val1 - val2
	}
	if operator == "*" {
		answer = val1 * val2
	}
	if operator == "%" {
		answer = val1 % val2
	}
	if operator == "/" {
		answer = val1 / val2
	}
	num := Itoa(answer)
	os.Stdout.WriteString(num + "\n")
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
	if isneg {
		result *= -1
	}
	return result
}

func Itoa(number int) string {
	isneg := false
	if number < 0 {
		isneg = true
		number *= -1
	}

	result := ""
	if number == 0 {
		return "0"
	}
	for number > 0 {
		digit := number % 10
		result = string(digit+'0') + result
		number /= 10
	}
	if isneg {
		return "-" + result
	}
	return result
}
